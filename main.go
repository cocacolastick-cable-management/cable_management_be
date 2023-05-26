package main

import (
	"github.com/cable_management/cable_management_be/src/app/routers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	routers.MountApiRouters(app)

	err := app.Listen(":8000")
	log.Fatal(err)
}
