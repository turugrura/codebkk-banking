package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turugrura/codebkk-banking/errs"
	"github.com/turugrura/codebkk-banking/logs"
	"github.com/turugrura/codebkk-banking/service"
)

type accountHandler struct {
	accService service.AccountService
}

func NewAccountHandler(accService service.AccountService) accountHandler {
	return accountHandler{accService: accService}
}

func (h accountHandler) NewAccount(c *fiber.Ctx) error {
	if c.Is("json") {
		return fiberError(errs.NewValidationError("request body incorect format"))
	}

	request := service.NewAccountRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		logs.Error(err)
		return fiberError(errs.NewValidationError("request body incorect format"))
	}

	customerId, err := c.ParamsInt("customerID")
	if err != nil {
		return fiberError(errs.NewValidationError("customerID should be integer"))
	}

	accRes, err := h.accService.NewAccount(customerId, request)
	if err != nil {
		logs.Error(err)
		return fiberError(err)
	}

	c.Status(fiber.StatusCreated).JSON(accRes)

	return nil
}

func (h accountHandler) GetAccounts(c *fiber.Ctx) error {
	customerId, err := c.ParamsInt("customerID")
	if err != nil {
		return fiberError(errs.NewValidationError("customerID should be integer"))
	}

	accResponses, err := h.accService.GetAccounts(customerId)
	if err != nil {
		logs.Error(err)
		return fiberError(err)
	}

	c.JSON(accResponses)

	return nil
}
