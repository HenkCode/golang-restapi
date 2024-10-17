package routes

import (
	"github.com/HenkCode/golang-restapi/controllers"
	"github.com/HenkCode/golang-restapi/middleware"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	
	e.GET("/", controllers.IndexPage)
	
	e.GET("/siswa", controllers.FetchAllSiswa, middleware.IsAuthenticated)
	e.POST("/siswa", controllers.StoreSiswa, middleware.IsAuthenticated)
	e.PUT("/siswa", controllers.UpdateSiswa, middleware.IsAuthenticated)
	e.DELETE("/siswa", controllers.DeleteSiswa, middleware.IsAuthenticated)
	
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.GET("/validation-struct", controllers.ValidationStruct)
	e.GET("/validation-variable", controllers.ValidationVariable)
	e.POST("/auth", controllers.CheckAuth)

	return e
}