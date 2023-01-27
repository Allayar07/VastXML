package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MessageError struct {
	Message string `json:"message"`
}

func ErrorMessage(c *fiber.Ctx, statuCode int, message string) error {
	logrus.Error(message)
	return c.Status(statuCode).JSON(MessageError{Message: message})
}
