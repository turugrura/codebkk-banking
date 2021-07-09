package service

import (
	"testing"

	"github.com/turugrura/codebkk-banking/repository"
)

func getCustomerService() CustomerService {
	repo := repository.NewCustomerRepositoryMock()

	return NewCustomerService(repo)
}

func getMockCustomers() []repository.Customer {
	return repository.MockCustomers
}

func TestGetCustomer(t *testing.T) {
	tests := []struct {
		customerId int
		expectName interface{}
	}{
		{customerId: 1, expectName: "one"},
		{customerId: 2, expectName: "two"},
	}

	custService := getCustomerService()
	for _, test := range tests {
		cusRes, err := custService.GetCustomer(test.customerId)
		if err != nil {
			t.Error("should not be an error")
		}

		if cusRes.Name != test.expectName {
			t.Errorf("expected %v | got %v", test.expectName, cusRes.Name)
		}
	}

	tests = []struct {
		customerId int
		expectName interface{}
	}{
		{customerId: 10},
		{customerId: 20},
	}
	for _, test := range tests {
		_, err := custService.GetCustomer(test.customerId)
		if err == nil {
			t.Error("should be a not found error")
		}
	}
}

func TestGetCustomers(t *testing.T) {
	custService := getCustomerService()

	mockCustomers := getMockCustomers()
	customers, err := custService.GetCustomers()
	if err != nil {
		t.Error("should not be an error")
	}

	for i, cust := range customers {
		gotName := cust.Name
		expectedName := mockCustomers[i].Name
		if cust.Name != expectedName {
			t.Errorf("expected %v | get %v", expectedName, gotName)
		}

		gotStatus := cust.Status
		expectedStatus := mockCustomers[i].Status
		if cust.Status != expectedStatus {
			t.Errorf("expected %v | get %v", expectedStatus, gotStatus)
		}
	}
}
