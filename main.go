package main

import (
	"fmt"
	"os"
	"test-product/app/database"
	"test-product/app/error"
	"test-product/app/routes"

	_ "test-product/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger" // swagger handler
	"github.com/joho/godotenv"
)

// @title TEST PRODUCT API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /

func InitFiber() *fiber.App {
	config := fiber.Config{
		ErrorHandler: error.CustomError,
	}
	return fiber.New(config)
}

func main() {
	godotenv.Load(".env")
	database.DBInit()

	app := InitFiber()
	app.Use(logger.New())
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	routes.RouteInit(app)

	port := os.Getenv("PORT")

	appErr := app.Listen(":" + port)

	if appErr != nil {
		fmt.Println("Server Init Error!")
		os.Exit(1)
	}
}
