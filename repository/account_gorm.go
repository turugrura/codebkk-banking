package repository

import "gorm.io/gorm"

type accountRepositoryGorm struct {
	db *gorm.DB
}

func NewAccountRepositoryGorm(db *gorm.DB) AccountRepository {
	return accountRepositoryGorm{db: db}
}

func (r accountRepositoryGorm) Create(acc Account) (*Account, error) {
	tx := r.db.Create(&acc)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &acc, nil
}

func (r accountRepositoryGorm) GetAll(customerID int) ([]Account, error) {
	accounts := []Account{}

	tx := r.db.Where("customer_id=?", customerID).Find(&accounts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return accounts, nil
}
