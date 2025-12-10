package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"resapiserver/database"

	"github.com/gofiber/fiber/v2"
)

func DeleteBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	// RETURNING id untuk cek data ada atau tidak
	row := database.DB.QueryRow("DELETE FROM books WHERE id=$1 RETURNING id", id)
	if row.Err() != nil {
		return row.Err()
	}

	var returnedId int
	err := row.Scan(&returnedId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]string{
				"message": "Not Found",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": "OK",
	})
}
