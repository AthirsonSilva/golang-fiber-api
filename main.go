package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sixfwa/fiber-api/database"
	"github.com/sixfwa/fiber-api/routes"
)

func SetupRoutes(app *fiber.App) {
	// Welcome route
	app.Get("/", welcome)

	// User routes
	app.Post("/api/v1/users", routes.CreateUser)
	app.Get("/api/v1/users", routes.GetAllUser)
	app.Get("/api/v1/users/:id", routes.GetUser)
	app.Put("/api/v1/users/:id", routes.UpdateUser)
	app.Delete("/api/v1/users/:id", routes.DeleteUser)

	// Product routes
	app.Get("/api/v1/products", routes.GetAllProduct)
	app.Post("/api/v1/products", routes.CreateProduct)
	app.Get("/api/v1/products/:id", routes.GetProduct)
	app.Put("/api/v1/products/:id", routes.UpdateProduct)
	app.Delete("/api/v1/products/:id", routes.DeleteProduct)

	// Order routes
	app.Post("/api/v1/orders", routes.CreateOrder)
	app.Get("/api/v1/orders", routes.GetAllOrder)
	app.Get("/api/v1/orders/:id", routes.GetOrder)
	app.Put("/api/v1/orders/:id", routes.UpdateOrder)
	app.Delete("/api/v1/orders/:id", routes.DeleteOrder)

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
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Fiber!")
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	database.Connect()

	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
