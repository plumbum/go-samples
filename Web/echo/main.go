package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
	"github.com/thoas/stats"
)

type Data struct {
	Id int
	Name string
	Tags []string
}

func main() {

	demoData := Data{
		Id: 5,
		Name: "User name",
		Tags: []string{"people", "customer", "developer"},
	}

	e := echo.New()
	e.SetDebug(true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := stats.New()
	e.Use(standard.WrapMiddleware(s.Handler))

	e.GET("/xml", func (c echo.Context) error {
		return c.XML(200, demoData)
	})

	e.GET("/json", func (c echo.Context) error {
		return c.JSON(200, demoData)
	})

	e.GET("/error", func (c echo.Context) error {
		return echo.NewHTTPError(500, "Error here")
	})

	e.Run(standard.New(":8888"))

}