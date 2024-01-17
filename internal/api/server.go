package api

import (
	"go-ecommerce-app/config"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {

	app := fiber.New()
	log.Println("Configured")

	app.Get("/health", HealthCheck)

	app.Listen(config.ServerPort)

}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "I am breathing!",
	})
}
