package main

import (
	"log"
	"net/http"
	"regexp"
	"resapiserver/controller"
	"resapiserver/database"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Cara 1
	// app := fiber.New()

	// Cara 2 error yang di return dari controller langsung sehingga dapat di handle di satu tempat
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			validationerrs, ok := err.(validator.ValidationErrors)

			if ok {
				errRes := make(map[string]string)
				for _, validationerr := range validationerrs {
					tag := validationerr.Tag()
					message := "unknown error"
					if tag == "required" {
						message = "Required"
					} else if tag == "min" {
						minimum := validationerr.Param()
						message = "Minimun should be " + minimum
					}

					// Cara 1
					// errRes[validationerr.Field()] = validationerr.Tag()

					// Cara 2
					// regex manipulasi string
					re := regexp.MustCompile("([a-z])([A-Z])")
					snakeCase := re.ReplaceAllString(validationerr.Field(), "${1}_${2}")

					errRes[strings.ToLower(snakeCase)] = message
				}

				return c.Status(http.StatusBadRequest).JSON(map[string]any{
					// Cara 1
					// "errors": err.Error(),

					// Cara 2
					"errors": errRes,
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(map[string]any{
				"message": "Internal Server Error",
			})
		},
	})

	// run db
	database.InitDb()

	// passing data jadi tidak perlu () untuk CreateBookController
	app.Post("api/books", controller.CreateBookController)

	app.Get("api/books/:id", controller.ShowBookController)

	app.Delete("api/books/:id", controller.DeleteBookController)

	app.Put("api/books/:id", controller.UpdateBookController)

	log.Fatal(app.Listen(":3000"))
}
