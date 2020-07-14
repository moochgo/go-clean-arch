package route

import (
	"go-clean-arch/delivery/http/controller"

	"github.com/gofiber/fiber"
)

// ControllerHandler ...
type ControllerHandler struct {
	CusHandler controller.CustomerHandler
}

// Run ...
func Run(appName string, port string, rHandler ControllerHandler) {
	handler := &ControllerHandler{
		CusHandler: rHandler.CusHandler,
	}
	app := fiber.New()
	app.Get("/", handler.CusHandler.GetCustomerByID)
	app.Listen("0.0.0.0:" + port)
}
