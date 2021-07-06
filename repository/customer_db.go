package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, status from customers where customer_id = @customer_id"
	err := r.db.Get(&customer, query, sql.Named("customer_id", id))
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
