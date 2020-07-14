package repository

import (
	"go-clean-arch/models"
)

// CustomerRepository : Customer repository ...
type CustomerRepository interface {
	GetByID(id int) (*models.Customer, error)
	FindAll() (customer []*models.Customer, err error)
}
