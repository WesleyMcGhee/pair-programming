package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Signin(c echo.Context) error {
	return c.String(http.StatusOK, "SignIn")
}

func Signup(c echo.Context) error {
	return c.String(http.StatusOK, "SignUp")
}