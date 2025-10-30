package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/login"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/register"
)

type UserHandler struct {
	registerUseCase register.UseCase
	loginUseCase    login.UseCase
}

func NewUserHandler(registerUC register.UseCase, loginUc login.UseCase) *UserHandler {
	return &UserHandler{registerUseCase: registerUC, loginUseCase: loginUc}
}

func (h *UserHandler) UserRoutes(router fiber.Router) {
	userRoutes := router.Group("/users")

	userRoutes.Post("/register", h.Register)
	userRoutes.Post("/login", h.Login)
}

func (h *UserHandler) Register(c fiber.Ctx) error {

	req, err := helper.BindAndValidate[register.DTO](c, helper.BindBody)

	if err != nil {
		return err
	}

	user, err := h.registerUseCase.Execute(register.DTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	req, err := helper.BindAndValidate[login.DTO](c, helper.BindBody)

	if err != nil {
		return err
	}

	user, err := h.loginUseCase.Execute(login.DTO{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(user)
}
