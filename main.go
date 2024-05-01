package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"log"
	"redmine-time-manager-notepad/docs"
)

type HTTPError struct {
	Message string
}

// Cli flags of working modes
var (
	mode     *int
	basePath = "/api/v1"
)

func init() {
	mode = flag.Int("mode", 0, "0 - PROD, 1 - DEV")
	flag.Parse()
}

// @title Redmine Notepad
// @description Redmine Notepad

// @contact.name George Zakharov
// @contact.email george_zakharov@mail.ru

// @schemes http
func main() {
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.BasePath = basePath

	// Generating Swagger API according to environment
	//switch *mode {
	//case 1:
	//	docs.SwaggerInfo.Host = configs.GetDevServiceHostName()
	//default:
	//	docs.SwaggerInfo.Host = configs.GetProdServiceHostName()
	//}

	app := fiber.New()
	app.Use(recover.New()) // Here we use middleware to recover from panics

	//route := app.Group(configs.BasePath)
	route := app.Group(basePath)

	route.Get("/swagger/*", swagger.HandlerDefault)

	route.Get("/", healthCheck)
	route.Get("/health", healthCheck)
	//route.Get("/version")

	log.Fatal(app.Listen(":" + "10000"))
}

// healthCheck godoc
// @Summary Show the status of this service.
// @Description Get the status of this service.
// @Tags rotator
// @Accept */*
// @Produce json
// @Success 200 {string} string
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router / [get]
// @Router /health [get]
func healthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("OK")
}
