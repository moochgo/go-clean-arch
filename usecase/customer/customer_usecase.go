package customer

import (
	"go-clean-arch/models"
	"go-clean-arch/repository"
	"go-clean-arch/usecase"
)

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

// NewCustomerUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewCustomerUsecase(customer repository.CustomerRepository) usecase.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customer,
	}
}

// GetByID ...
func (cus *customerUsecase) GetByID(id int) (*models.Customer, error) {
	customer, err := cus.customerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// FindAll ...
func (cus *customerUsecase) FindAll() (customer []*models.Customer, err error) {
	customers, err := cus.customerRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}
