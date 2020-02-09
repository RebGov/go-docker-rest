package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func goWithDocker(c echo.Context) error {
	return c.JSON(http.StatusOK, "Becci Go with Docker Container using Echo")
}

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) //default
	log.Level = logrus.DebugLevel
}

func main() {
	fmt.Println("Hello Becci's World!")
	fmt.Println("Main function :")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//echo Routes
	e.GET("/main", goWithDocker)
	e.Logger.Fatal(e.Start(":8080"))
}
