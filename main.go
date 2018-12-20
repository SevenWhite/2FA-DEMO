package main

import (
	"bytes"
	"fmt"
	"github.com/SevenWhite/2FA-DEMO/store"
	"github.com/labstack/echo"
	"github.com/pquerna/otp/totp"
	"image/jpeg"
	"net/http"
	"os"
)

type ValidationInput struct {
	Code string `json:"code"`
}

func main() {
	store.AddUser("r.ponomarenko@digiteum.com", "123")

	e := echo.New()
	e.Static("/public", "public")
	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, store.Users())
	})
	e.GET("/users/:email/qr", func(c echo.Context) error {
		email := c.Param("email")
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "MyCompany.com",
			AccountName: email,
		})
		if err != nil {
			return err
		}
		store.UpdateUserSecret(email, key.Secret())
		// Convert TOTP key into a PNG
		var buf bytes.Buffer
		img, err := key.Image(200, 200)
		if err != nil {
			panic(err)
		}

		if err := jpeg.Encode(&buf, img, nil); err != nil {
			return err
		}

		return c.Stream(http.StatusOK, "image/jpeg", &buf)
	})
	e.POST("/users/:email/validate", func(c echo.Context) error {
		email := c.Param("email")
		usr, err := store.FindUserByEmail(email)
		if err != nil {
			return err
		}

		input := ValidationInput{}
		if err := c.Bind(&input); err != nil {
			return err
		}

		result := totp.Validate(input.Code, usr.Secret)
		return c.JSON(http.StatusOK, result)
	})
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
