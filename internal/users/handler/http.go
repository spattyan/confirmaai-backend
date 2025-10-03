package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/users/domain"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/login"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/register"
)

type UserHandler struct {
	registerUseCase register.UseCase
	loginUseCase    login.UseCase
}

func NewUserHandler(repository domain.Repository) *UserHandler {
	return &UserHandler{registerUseCase: register.NewUseCase(repository), loginUseCase: login.NewUseCase(repository)}
}

func (h *UserHandler) UserRoutes(router fiber.Router) {
	userRoutes := router.Group("/users")

	userRoutes.Post("/register", h.Register)
	userRoutes.Post("/login", h.Login)
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	var req register.Request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	validate, err := helper.Validate(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(validate)
	}

	user, err := h.registerUseCase.Execute(register.DTO{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *UserHandler) Login(c fiber.Ctx) error {
	var req login.Request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	validate, err := helper.Validate(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(validate)
	}

	user, err := h.loginUseCase.Execute(login.DTO{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(user)
}
