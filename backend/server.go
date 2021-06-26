package main

import (
	"net/http"

	"server/controller"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/test", controller.TestAuth)

	e.Logger.Fatal(e.Start(":1323"))
}
