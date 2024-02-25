package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateStudent(c echo.Context) error {
	return c.String(http.StatusOK, "CreateStudent")	
}