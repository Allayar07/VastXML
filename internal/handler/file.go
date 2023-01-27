package handler

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

const Directory = "uploads/"

func (h *Handler) OpenFile(c *fiber.Ctx) error {
	fileName := c.Params("any")

	fileBytes, err := os.ReadFile(Directory + fileName)
	if err != nil {
		return ErrorMessage(c, 500, err.Error())
	}

	_, err = c.Write(fileBytes)

	return err
}
