package main

import (
	"fmt"
	"github.com/SevenWhite/2FA-DEMO/store"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func main() {
	store.AddUser("r.ponomarenko@digiteum.com", "123")

	e := echo.New()
	e.Static("/public", "public")
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, store.Users())
	})
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
