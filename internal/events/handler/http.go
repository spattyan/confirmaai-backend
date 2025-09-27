package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/create"
)

type EventHandler struct {
	createUseCase create.UseCase
}

func NewEventHandler(repository domain.Repository) *EventHandler {
	return &EventHandler{createUseCase: create.NewUseCase(repository)}
}

func (h *EventHandler) EventRoutes(router fiber.Router) {
	router.Post("/events/new", h.Create)
}

func (h *EventHandler) Create(c fiber.Ctx) error {
	var req create.Request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	event, err := h.createUseCase.Execute(create.DTO{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		DateAndTime:      req.DateAndTime,
		ParticipantLimit: req.ParticipantLimit,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "An internal server error occurred"})
	}

	response := create.Response{
		ID: event.ID.String(),
	}

	return c.Status(http.StatusCreated).JSON(response)
}
