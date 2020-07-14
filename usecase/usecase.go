package usecase

import (
	"go-clean-arch/models"
)

// CustomerUsecase ...
type CustomerUsecase interface {
	GetByID(id int) (*models.Customer, error)
	FindAll() (customer []*models.Customer, err error)
}
