package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turugrura/codebkk-banking/errs"
)

func fiberError(err error) error {
	switch v := err.(type) {
	case errs.AppError:
		return fiber.NewError(v.Code, v.Message)
	case error:
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return err
}
