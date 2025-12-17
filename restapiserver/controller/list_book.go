package controller

import (
	"net/http"
	"resapiserver/database"
	"resapiserver/dto"

	"github.com/gofiber/fiber/v2"
)

func ListBookController(c *fiber.Ctx) error {
	// Query karena multiple baris
	rows, err := database.DB.Query("SELECT id, title, description, isbn, author, genre, publish_date FROM books")
	if err != nil {
		return err
	}

	var res []dto.ListBookResponse

	for rows.Next() {
		var book dto.ListBookResponse
		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Description,
			&book.ISBN,
			&book.Author,
			&book.Genre,
			&book.PublishDate,
		)
		if err != nil {
			return err
		}

		res = append(res, book)
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": res,
	})

}
