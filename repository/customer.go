package repository

type Customer struct {
	CustomerID int    `db:"customer_id"`
	Name       string `db:"name"`
	Status     int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
