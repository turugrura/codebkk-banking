package service

import (
	"database/sql"
	"errors"
	"log"

	"github.com/turugrura/codebkk-banking/repository"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: repo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			Name:   customer.Name,
			Status: customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}

		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		Name:   customer.Name,
		Status: customer.Status,
	}

	return &custResponse, nil
}
