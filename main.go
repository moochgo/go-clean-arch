package main

import (
	"go-clean-arch/delivery/http/controller"
	"go-clean-arch/delivery/http/route"
	"go-clean-arch/repository/mysql"
	"os"
	"strings"

	_customerRepo "go-clean-arch/repository/mysql"
	_customerUsecase "go-clean-arch/usecase/customer"

	"github.com/subosito/gotenv"
)

func init() {
	if err := gotenv.Load(); err != nil {
		if strings.Contains(err.Error(), "no such file") {
			gotenv.Load()
		}
	}
}

func main() {
	db := mysql.DBs{}
	db.CustomerConnection()
	defer db.Customer.Close()
	cusRepo := _customerRepo.NewMysqlCustomerRepository(db.Customer)
	cusUsecase := _customerUsecase.NewCustomerUsecase(cusRepo)
	cusController := controller.CustomerHandler{
		CusUsecase: cusUsecase,
	}
	handler := &route.ControllerHandler{
		CusHandler: cusController,
	}
	route.Run(os.Getenv("APP_NAME"), os.Getenv("APP_PORT"), *handler)
}
