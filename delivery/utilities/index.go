package utilities

import (
	"fmt"
	"os"

	"go-clean-arch/delivery/utilities/responses"

	"github.com/gofiber/fiber"
)

// ResponseRender ...
func ResponseRender(types string, ctx *fiber.Ctx, statusCode int, v interface{}) {
	if types == "error" {
		responseData := responses.Standard{
			Status:       "error",
			ErrorMessage: fmt.Sprintf("%v", v),
			Data:         nil,
			RequestParam: "",
			Next:         "",
			Version:      responses.Version{Code: os.Getenv("APP_VERSION_CODE"), Name: os.Getenv("APP_VERSION")},
		}
		ctx.Status(statusCode).JSON(responseData)
	} else {
		responseData := responses.Standard{
			Status:       "success",
			ErrorMessage: "",
			Data:         v,
			RequestParam: "",
			Next:         "",
			Version:      responses.Version{Code: os.Getenv("APP_VERSION_CODE"), Name: os.Getenv("APP_VERSION")},
		}
		ctx.Status(statusCode).JSON(responseData)
	}
}
