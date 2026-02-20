package middleware

import (
	"net/http"
	"strings"

	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"github.com/labstack/echo/v4"
)

const JwtClaimsKey = "jwt_claims"

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}

			claims, err := security.ValidateJWT(parts[1])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}

			c.Set(JwtClaimsKey, claims)
			return next(c)
		}
	}
}

func GetClaims(c echo.Context) (*security.JwtClaims, bool) {
	claims, ok := c.Get(JwtClaimsKey).(*security.JwtClaims)
	return claims, ok
}
