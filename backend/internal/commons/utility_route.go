package commons

import (
	"net/http"

	"github.com/iacopoGhilardi/amILate/internal/db"
	"github.com/labstack/echo/v4"
)

func InitUtilityRoute(e *echo.Group) {

	e.GET("/health", func(c echo.Context) error {
		if err := db.Ping(); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status": "Db unreachable",
				"error":  err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

}
