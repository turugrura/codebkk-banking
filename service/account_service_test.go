package service

import (
	"testing"

	"github.com/turugrura/codebkk-banking/repository"
)

func getAccountService() AccountService {
	repo := repository.NewAccountRepositoryMock()

	return NewAccountService(repo)
}

func getMockAccounts() []repository.Account {
	return repository.MockAccounts
}

func TestGetAccounts(t *testing.T) {
	custService := getAccountService()
	customerID := 1

	mockAccounts := getMockAccounts()
	// expectedAcc := []AccountResponse{}
	expectedAccID := make(map[int]int)
	for _, acc := range mockAccounts {
		if acc.CustomerID == customerID {
			// expectedAcc = append(expectedAcc, AccountResponse{
			// 	AccountID: acc.AccountID,
			// 	OpeningDate: acc.OpeningDate,
			// 	AccountType: acc.AccountType,
			// 	Amount: acc.Amount,
			// 	Status: acc.Status,
			// })
			expectedAccID[acc.AccountID] = acc.AccountID
		}
	}

	accounts, err := custService.GetAccounts(customerID)
	if err != nil {
		t.Error("should not be an error")
	}

	if len(expectedAccID) != len(accounts) {
		t.Errorf("expected %v | got %v", len(expectedAccID), len(accounts))
	}

	for _, acc := range accounts {
		if _, ok := expectedAccID[acc.AccountID]; !ok {
			t.Errorf("not found accountID = %v", acc.AccountID)
		}
	}
}
