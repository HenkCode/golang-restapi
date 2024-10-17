package controllers

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Siswa struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"required,lte=20,gte=7"`
}

func ValidationVariable(c echo.Context) error {
	v := validator.New()

	email := "henk@gmail.com"
	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}

func ValidationStruct(c echo.Context) error {
	v := validator.New()

	siswa := Siswa{
		Nama: "henk",
		Email: "henk@gmail.com",
		Alamat: "kalikidang",
		Umur: 21,
	}
	
	err := v.Struct(siswa)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success"})
}

