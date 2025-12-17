package controller

import (
	"net/http"
	"resapiserver/database"
	"resapiserver/dto"

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
		return err
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
