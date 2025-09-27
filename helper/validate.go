package helper

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func Validate(s interface{}) (fiber.Map, error) {
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(s); err != nil {
		var errors []string
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				errors = append(errors, fmt.Sprintf("%s (%s): %s", fieldErr.Field(), fieldErr.Type(), fieldErr.Tag()))
			}
			return fiber.Map{"errors": errors}, err
		}
		return fiber.Map{"error": err.Error()}, err
	}
	return fiber.Map{"status": http.StatusOK}, nil
}
