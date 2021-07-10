package repository

type accountRepositoryMock struct {
	accounts *[]Account
}

var MockAccounts = []Account{
	{AccountID: 1, CustomerID: 1, Amount: 1000, Status: 1, AccountType: "checking"},
	{AccountID: 2, CustomerID: 1, Amount: 1000, Status: 1, AccountType: "saving"},
	{AccountID: 3, CustomerID: 2, Amount: 1000, Status: 1, AccountType: "saving"},
	{AccountID: 4, CustomerID: 2, Amount: 1000, Status: 1, AccountType: "saving"},
	{AccountID: 5, CustomerID: 2, Amount: 1000, Status: 1, AccountType: "checking"},
}

func NewAccountRepositoryMock() AccountRepository {
	return accountRepositoryMock{accounts: &MockAccounts}
}

func (r accountRepositoryMock) Create(acc Account) (*Account, error) {
	acc.AccountID = len(*r.accounts) + 1
	*r.accounts = append(*r.accounts, acc)

	return &acc, nil
}

func (r accountRepositoryMock) GetAll(customerID int) ([]Account, error) {
	accounts := []Account{}

	for _, acc := range *r.accounts {
		if acc.CustomerID == customerID {
			accounts = append(accounts, acc)
		}
	}

	return accounts, nil
}
