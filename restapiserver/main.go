package main

import (
	"log"
	"resapiserver/controller"
	"resapiserver/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// run db
	database.InitDb()

	// passing data jadi tidak perlu () untuk CreateBookController
	app.Post("api/books", controller.CreateBookController)

	app.Get("api/books/:id", controller.ShowBookController)

	log.Fatal(app.Listen(":3000"))
}
