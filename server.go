package main

import (
	"context"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func main() {
	e := echo.New()

	ctx := context.Background()
	serviceAccount := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		fmt.Println("error:", err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		fmt.Println("error:", err)
	}

	e.POST("/uid/:uid/claims/admin", func(c echo.Context) error {
		uid := c.Param("uid")

		claims := &map[string]interface{}{
			"https://hasura.io/jwt/claims": map[string]interface{}{
				"x-hasura-default-role":  "admin",
				"x-hasura-allowed-roles": []string{"admin"},
			},
		}

		err := auth.SetCustomUserClaims(ctx, uid, *claims)
		if err != nil {
			fmt.Println("set claim error:", err)
		}

		return c.String(http.StatusOK, "set claim success")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8888"))
}
