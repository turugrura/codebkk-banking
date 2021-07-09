package repository

import "errors"

type customerRepositoryMock struct {
	customers []Customer
}

var MockCustomers = []Customer{
	{CustomerID: 1, Name: "one", Status: 1},
	{CustomerID: 2, Name: "two", Status: 0},
	{CustomerID: 3, Name: "three", Status: 1},
	{CustomerID: 4, Name: "four", Status: 0},
	{CustomerID: 5, Name: "five", Status: 1},
}

func NewCustomerRepositoryMock() CustomerRepository {
	return customerRepositoryMock{customers: MockCustomers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return r.customers, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range r.customers {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
