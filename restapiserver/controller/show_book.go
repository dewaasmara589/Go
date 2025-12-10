package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"resapiserver/database"
	"resapiserver/dto"

	"github.com/gofiber/fiber/v2"
)

func ShowBookController(c *fiber.Ctx) error {
	// key sesuai param get yang ada di main.go
	id := c.Params("id")

	row := database.DB.QueryRow("SELECT id, title, description, isbn, author, genre, publish_date FROM books WHERE id=$1", id)

	if row.Err() != nil {
		return row.Err()
	}

	var res dto.ShowBookResponse
	err := row.Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.ISBN,
		&res.Author,
		&res.Genre,
		&res.PublishDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"message": "Not Found",
			})
		}
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": res,
	})
}
