package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/turugrura/codebkk-banking/errs"
	"github.com/turugrura/codebkk-banking/service"
)

type customerHandler struct {
	custService service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custService: custSrv}
}

func (h customerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.custService.GetCustomers()
	if err != nil {
		return fiberError(err)
	}

	c.JSON(customers)

	return nil
}

func (h customerHandler) GetCustomer(c *fiber.Ctx) error {
	customerId, err := c.ParamsInt("customerID")
	if err != nil {
		return fiberError(errs.NewValidationError("customerID should be integer"))
	}

	customer, err := h.custService.GetCustomer(customerId)
	if err != nil {
		return fiberError(err)
	}

	c.JSON(customer)

	return nil
}
