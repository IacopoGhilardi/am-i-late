package utils

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func ParseIDParam(c echo.Context, name string) (uint, error) {
	var id uint
	_, err := fmt.Sscan(c.Param(name), &id)
	return id, err
}
