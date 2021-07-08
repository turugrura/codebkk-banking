package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/turugrura/codebkk-banking/errs"
)

func handleError(c *fiber.Ctx, err error) error {
	switch v := err.(type) {
	case errs.AppError:
		c.Response().Header.SetStatusCode(v.Code)
	case error:
		c.Response().Header.SetStatusCode(http.StatusInternalServerError)
	}

	return err
}
