package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/iacopoGhilardi/amILate/internal/middleware"
	"github.com/iacopoGhilardi/amILate/internal/utils/security"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type JWTMiddlewareTestSuite struct {
	suite.Suite
	echo *echo.Echo
}

func (s *JWTMiddlewareTestSuite) SetupTest() {
	s.echo = echo.New()
	security.SetJWTSecret("test-secret")
}

func (s *JWTMiddlewareTestSuite) newContext(token string) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	rec := httptest.NewRecorder()
	return s.echo.NewContext(req, rec), rec
}

func (s *JWTMiddlewareTestSuite) applyMiddleware(c echo.Context) error {
	h := middleware.JWTMiddleware()(func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	return h(c)
}

func (s *JWTMiddlewareTestSuite) Test_should_return_unauthorized_when_authorization_header_is_missing() {
	c, _ := s.newContext("")

	err := s.applyMiddleware(c)

	he, ok := err.(*echo.HTTPError)
	s.True(ok)
	s.Equal(http.StatusUnauthorized, he.Code)
	s.Equal("missing authorization header", he.Message)
}

func (s *JWTMiddlewareTestSuite) Test_should_return_unauthorized_when_authorization_header_has_invalid_format() {
	c, _ := s.newContext("InvalidTokenWithoutBearer")

	err := s.applyMiddleware(c)

	he, ok := err.(*echo.HTTPError)
	s.True(ok)
	s.Equal(http.StatusUnauthorized, he.Code)
	s.Equal("invalid authorization header format", he.Message)
}

func (s *JWTMiddlewareTestSuite) Test_should_return_unauthorized_when_token_is_invalid() {
	c, _ := s.newContext("Bearer tokenmalformato")

	err := s.applyMiddleware(c)

	he, ok := err.(*echo.HTTPError)
	s.True(ok)
	s.Equal(http.StatusUnauthorized, he.Code)
	s.Equal("invalid or expired token", he.Message)
}

func (s *JWTMiddlewareTestSuite) Test_should_return_unauthorized_when_token_is_signed_with_wrong_secret() {
	security.SetJWTSecret("wrong-secret")
	token, _ := security.GenerateJWT(uuid.New(), "test@example.com")

	security.SetJWTSecret("test-secret")
	c, _ := s.newContext("Bearer " + token)

	err := s.applyMiddleware(c)

	he, ok := err.(*echo.HTTPError)
	s.True(ok)
	s.Equal(http.StatusUnauthorized, he.Code)
}

func (s *JWTMiddlewareTestSuite) Test_should_pass_when_token_is_valid() {
	userId := uuid.New()
	token, err := security.GenerateJWT(userId, "test@example.com")
	s.NoError(err)

	c, rec := s.newContext("Bearer " + token)

	h := middleware.JWTMiddleware()(func(c echo.Context) error {
		claims, ok := middleware.GetClaims(c)
		s.True(ok)
		s.Equal(userId, claims.UserId)
		s.Equal("test@example.com", claims.Email)
		return c.String(http.StatusOK, "ok")
	})

	err = h(c)
	s.NoError(err)
	s.Equal(http.StatusOK, rec.Code)
}

func TestJWTMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(JWTMiddlewareTestSuite))
}
