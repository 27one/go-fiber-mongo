package main

import (
	"log"

	"github.com/27one/go-fiber-mongo/pkg/config"
	"github.com/27one/go-fiber-mongo/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", controllers.CreateEmployee)
	app.Post("/employee", controllers.CreateEmployee)
	app.Put("/employee/:id", controllers.UpdateEmployee)
	app.Delete("/employee/:id", controllers.DeleteEmployee)

	log.Fatal(app.Listen(":3000"))
}
