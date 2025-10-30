package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validate.RegisterValidation("datetime", validateRFC3339)
}

func validateRFC3339(fl validator.FieldLevel) bool {
	dateString := fl.Field().String()
	_, err := time.Parse(time.RFC3339, dateString)
	return err == nil
}

func getErrorMessage(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return "this field is required."
	case "uuid":
		return "must be a valid UUID."
	case "email":
		return "must be a valid email address."
	case "min":
		return fmt.Sprintf("must to be longer than %s caracteres.", fieldErr.Param())
	case "max":
		return fmt.Sprintf("must to be lower than %s caracteres.", fieldErr.Param())
	case "gte":
		return fmt.Sprintf("the number needs to be greater than %s.", fieldErr.Param())
	case "lte":
		return fmt.Sprintf("the number needs to be lower than %s.", fieldErr.Param())
	case "datetime":
		return "datetime must be in format RFC3339 (ex: 2025-10-29T21:15:00-03:00)"
	default:
		return fmt.Sprintf("validation failed: %s", fieldErr.Tag())
	}
}

type BindSource int

const (
	BindBody BindSource = iota
	BindURI
	BindQuery
)

type ValidationError struct {
	Errors map[string]string
}

func (e *ValidationError) Error() string {
	bytes, _ := json.Marshal(e.Errors)
	return string(bytes)
}

func Validate(data interface{}) error {
	if err := validate.Struct(data); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errors := make(map[string]string)

			var structType reflect.Type
			if s, ok := data.(reflect.Value); ok {
				structType = s.Type()
			} else {
				structType = reflect.TypeOf(data)
			}
			if structType.Kind() == reflect.Ptr {
				structType = structType.Elem()
			}

			for _, fieldErr := range validationErrors {
				fieldName := fieldErr.StructField()
				field, _ := structType.FieldByName(fieldName)

				messageTag := field.Tag.Get("message")

				jsonFieldName := fieldErr.Field()

				if messageTag != "" {
					errors[jsonFieldName] = messageTag
				} else {
					errors[jsonFieldName] = getErrorMessage(fieldErr)
				}
			}

			return &ValidationError{Errors: errors}
		}
		return fiber.NewError(http.StatusInternalServerError, "Failed during validation")
	}
	return nil
}

func BindAndValidate[T any](c fiber.Ctx, source BindSource) (*T, error) {
	req := new(T)
	var bindErr error

	switch source {
	case BindBody:
		bindErr = c.Bind().Body(req)
	case BindURI:
		bindErr = c.Bind().URI(req)
	case BindQuery:
		bindErr = c.Bind().Query(req)
	default:
		return nil, fiber.NewError(http.StatusInternalServerError, "Invalid bind source specified")
	}

	if bindErr != nil {
		return nil, fiber.NewError(http.StatusBadRequest, "Cannot parse request")
	}

	if err := Validate(req); err != nil {
		return nil, err
	}

	return req, nil
}

func HandleError(c fiber.Ctx, err error) error {
	if vErr, ok := err.(*ValidationError); ok {
		return c.Status(http.StatusBadRequest).JSON(vErr.Errors)
	}

	if fErr, ok := err.(*fiber.Error); ok {
		return c.Status(fErr.Code).JSON(fiber.Map{"error": fErr.Message})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}
