package mysql

import (
	"go-clean-arch/models"
	"go-clean-arch/repository"

	"github.com/jinzhu/gorm"
)

// type mysqlCustomerRepository struct {
// 	Conn *gorm.DB
// }

// NewMysqlCustomerRepository will create an object that represent the article.Repository interface
func NewMysqlCustomerRepository(Conn *gorm.DB) repository.CustomerRepository {
	return &DBs{
		Customer: Conn,
	}
}

// GetByID ...
func (m *DBs) GetByID(id int) (*models.Customer, error) {
	customer := models.Customer{}
	err := m.Customer.Where("id = ? and deleted = 0", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// FindAll ...
func (m *DBs) FindAll() ([]*models.Customer, error) {
	customers := []*models.Customer{}
	err := m.Customer.Where("deleted = 0").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}
