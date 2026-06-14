package main

import (
	"github.com/gofiber/fiber/v2"
)

func errorHandler(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(fiber.Map{
		"message": err.Error(),
		"status":  code,
		"data":    nil,
		"error":   err.Error(),
	})
}
