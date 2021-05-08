package controllers

import (
	"golang-boilerplate/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFrameworkList(c echo.Context) error {
	var frameworks []models.Framework

	// frameworks = []models.Framework{
	// 	{
	// 		Name: "Gin",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Echo",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Revel",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Buffalo",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Iris",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Beego",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Fiber",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Chi",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Atreugo",
	// 		Type: "Web",
	// 	},
	// 	{
	// 		Name: "Gearbox",
	// 		Type: "Web",
	// 	},
	// }

	err := models.GetFrameworkList(&frameworks)
	if err != http.StatusOK {
		return c.JSON(err, http.StatusText(err))
	}

	return c.JSON(http.StatusOK, frameworks)
}

func GetFramework(c echo.Context) error {
	name := c.Param("name")

	var framework models.Framework
	err := models.GetFramework(&framework, name)
	if err != http.StatusOK {
		return c.JSON(err, http.StatusText(err))
	}

	return c.JSON(http.StatusOK, framework)
}

func CreateFramework(c echo.Context) error {
	name := c.FormValue("name")
	ftype := c.FormValue("type")

	framework := models.Framework{
		Name: name,
		Type: ftype,
	}

	err := models.CreateFramework(&framework)
	if err != http.StatusOK {
		return c.JSON(err, http.StatusText(err))
	}

	return c.JSON(http.StatusOK, "Framework created")
}

func ChangeFramework(c echo.Context) error {
	name := c.Param("name")
	ftype := c.FormValue("type")

	framework := models.Framework{
		Name: name,
		Type: ftype,
	}

	err := models.ChangeFramework(&framework, name)
	if err != http.StatusOK {
		return c.JSON(err, http.StatusText(err))
	}

	return c.JSON(http.StatusOK, "Framework "+name+" is changed")
}

func DeleteFramework(c echo.Context) error {
	name := c.Param("name")

	err := models.DeleteFramework(name)
	if err != http.StatusOK {
		return c.JSON(err, http.StatusText(err))
	}

	return c.JSON(http.StatusOK, "Framework "+name+" was deleted")
}
