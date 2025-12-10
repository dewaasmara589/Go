package controller

import (
	"net/http"
	"regexp"
	"resapiserver/database"
	"resapiserver/dto"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateBookController(c *fiber.Ctx) error {
	var request dto.CreateBookRequest
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	// validation
	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		validationerrs := err.(validator.ValidationErrors)

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

	// Exec cocok untuk yang tidak ada return valuenya
	// Query(banyak Data) QueryRow(1 Data saja) cocok untuk yang ada return valuenya

	// Cara 1
	// _, err = database.DB.Exec(

	// Cara 2
	row := database.DB.QueryRow(
		// $ cara posgresql mengihindari sql injection
		// RETURNING khusus hanya di posgresql
		"INSERT INTO books (title, description, author, genre, isbn, stock, publish_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		request.Title,
		request.Description,
		request.ISBN,
		request.Author,
		request.Genre,
		request.Stock,
		request.PublishDate,
	)
	if row.Err() != nil {
		return row.Err()
	}

	var id int
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": map[string]int{
			"id": id,
		},
	})
}
