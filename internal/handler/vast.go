package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

func (h *Handler) generateVast(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	
// 	change this to id := c.ParamsInt("id)
	
	if err != nil {
		return ErrorMessage(c, http.StatusBadRequest, "invalid id param")
	}
	if err = h.service.Vast.GenerateVast(id); err != nil {
		return ErrorMessage(c, http.StatusInternalServerError, err.Error())
	}
	return c.Status(200).JSON("OK")
	
// 	I think it will look pretty like this
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Result":"OK"})
}
