package main

import (
	"errors"
	"fmt"
	"golang-boilerplate/controllers"
	"golang-boilerplate/libs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	libs.Db = libs.MariaDBInit()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/frameworks", controllers.GetFrameworkList)
	e.GET("/framework/:name", controllers.GetFramework)
	e.POST("/framework", controllers.CreateFramework)
	e.PATCH("/framework/:name", controllers.ChangeFramework)
	e.DELETE("/framework/:name", controllers.DeleteFramework)

	e.Logger.Fatal(e.Start(":1323"))
}

func Greeting(name string) error {
	if name == "" {
		return errors.New("Insert your name")
	}

	fmt.Println("Hello " + name)
	return nil
}
