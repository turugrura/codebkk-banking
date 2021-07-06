package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepository {
	return accountRepository{db: db}
}

func (r accountRepository) Create(acc Account) (*Account, error) {
	query := "insert into accounts (customer_id, opening_date, account_type, amount, status) values (@cus_id, @open_date, @acc_type, @amt, @status)"
	result, err := r.db.Exec(query,
		sql.Named("cus_id", acc.CustomerID),
		sql.Named("open_date", acc.OpeningDate),
		sql.Named("acc_type", acc.AccountType),
		sql.Named("amt", acc.Amount),
		sql.Named("status", acc.Status),
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)

	return &acc, nil
}

func (r accountRepository) GetAll(customerID int) ([]Account, error) {
	query := "select account_id, customer_id, opening_date, account_type, amount, status from accounts where customer_id = @cust_id"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, sql.Named("cust_id", customerID))
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
