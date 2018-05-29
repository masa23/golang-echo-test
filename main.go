package main

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Data struct {
	Name string `json:"name"`
}

var list []*Data

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/data", getData)
	e.POST("/data", postData)
	e.Start(":8080")
}

func getData(c echo.Context) error {
	return c.JSON(http.StatusOK, list)
}

func postData(c echo.Context) error {
	data := new(Data)
	if err := c.Bind(data); err != nil {
		return err
	}
	list = append(list, data)
	pp.Println(list)

	return c.JSON(http.StatusCreated, data)
}
