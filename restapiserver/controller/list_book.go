package controller

import (
	"fmt"
	"net/http"
	"resapiserver/database"
	"resapiserver/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ListBookController(c *fiber.Ctx) error {
	var request dto.ListBookRequest
	err := c.QueryParser(&request)
	if err != nil {
		return err
	}

	query := "SELECT id, title, description, isbn, author, genre, publish_date FROM books"
	var args []any
	if request.Search != "" {
		// Cara 1
		// query += "WHERE LOWER(title) LIKE $1"

		// Cara 2 Multiple column
		query += " WHERE LOWER(title) LIKE $1 OR LOWER(description) LIKE $1 OR LOWER(isbn) LIKE $1 OR LOWER(author) LIKE $1 OR LOWER(genre) LIKE $1"

		// untuk cetak % perlu escape jadi %%
		filter := fmt.Sprintf("%%%s%%", strings.ToLower(request.Search))
		args = append(args, filter)
	}

	// Query karena multiple baris
	// args... adalah spread operator
	rows, err := database.DB.Query(query, args...)
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
