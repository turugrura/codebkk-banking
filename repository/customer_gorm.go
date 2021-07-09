package repository

import (
	"gorm.io/gorm"
)

type customerRepositoryGorm struct {
	db *gorm.DB
}

func NewCustomerRepositoryGorm(db *gorm.DB) CustomerRepository {
	return customerRepositoryGorm{db: db}
}

func (r customerRepositoryGorm) GetAll() ([]Customer, error) {
	customers := []Customer{}

	tx := r.db.Find(&customers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return customers, nil
}

func (r customerRepositoryGorm) GetById(id int) (*Customer, error) {
	customer := Customer{}

	tx := r.db.First(&customer, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &customer, nil
}
