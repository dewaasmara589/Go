package controller

import (
	"database/sql"
	"errors"
	"net/http"
	"resapiserver/database"
	"resapiserver/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UpdateBookController(c *fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateBookRequest
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return err
	}

	row := database.DB.QueryRow(
		// $ cara posgresql mengihindari sql injection
		// RETURNING khusus hanya di posgresql
		"UPDATE books SET title=$1, description=$2, author=$3, genre=$4, isbn=$5, stock=$6, publish_date=$7 WHERE id=$8 RETURNING id",
		req.Title,
		req.Description,
		req.ISBN,
		req.Author,
		req.Genre,
		req.Stock,
		req.PublishDate,
		id,
	)
	if row.Err() != nil {
		return row.Err()
	}

	var returnedId int
	err = row.Scan(&returnedId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(http.StatusNotFound).JSON(map[string]any{
				"message": "Not Found",
			})
		}
		return err
	}

	return c.Status(http.StatusOK).JSON(map[string]any{
		"data": map[string]any{
			"id": returnedId,
		},
	})
}
