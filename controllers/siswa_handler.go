package controllers

import (
	"net/http"
	"strconv"
	"github.com/HenkCode/golang-restapi/models"
	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

func FetchAllSiswa(c echo.Context) error {
	result, err := models.FetchSiswa()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSiswa(c echo.Context) error {
	nama := c.FormValue("nama") 
	alamat := c.FormValue("alamat") 
	nohp := c.FormValue("nohp") 
	
	result, err := models.StoreSiswa(nama, alamat, nohp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	return c.JSON(http.StatusOK, result)
}

func UpdateSiswa(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama") 
	alamat := c.FormValue("alamat") 
	nohp := c.FormValue("nohp")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateSiswa(conv_id, nama, alamat, nohp)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteSiswa(c echo.Context) error {
	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	
	result, err := models.DeleteSiswa(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}
	
	return c.JSON(http.StatusOK, result)
}