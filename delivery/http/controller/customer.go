package controller

import (
	"go-clean-arch/usecase"
	"log"
	"net/http"
	"strconv"

	srv "go-clean-arch/delivery/utilities"

	"github.com/gofiber/fiber"
)

// CustomerHandler ...
type CustomerHandler struct {
	CusUsecase usecase.CustomerUsecase
}

// GetCustomerByID ...
func (cus *CustomerHandler) GetCustomerByID(c *fiber.Ctx) {
	cusID, _ := strconv.Atoi(c.Query("cus_id"))
	customer, _ := cus.CusUsecase.GetByID(cusID)
	log.Println("customer ", customer)
	if customer == nil {
		srv.ResponseRender("success", c, http.StatusOK, "Data tidak ditemukan")
		return
	}
	srv.ResponseRender("success", c, http.StatusOK, customer)
	return
}

// FindAll ...
func (cus *CustomerHandler) FindAll(c *fiber.Ctx) {
	customers, _ := cus.CusUsecase.FindAll()
	if customers == nil {
		srv.ResponseRender("success", c, http.StatusOK, "Data tidak ditemukan")
		return
	}
	srv.ResponseRender("success", c, http.StatusOK, customers)
	return
}
