package service

type CustomerResponse struct {
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
