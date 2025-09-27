package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/create"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/list"
)

type EventHandler struct {
	createUseCase create.UseCase
	listUseCase   list.UseCase
}

func NewEventHandler(repository domain.Repository) *EventHandler {
	return &EventHandler{createUseCase: create.NewUseCase(repository), listUseCase: list.NewUseCase(repository)}
}

func (h *EventHandler) EventRoutes(router fiber.Router) {
	router.Post("/events/new", h.Create)
	router.Get("/events", h.List)
}

func (h *EventHandler) List(c fiber.Ctx) error {
	events, err := h.listUseCase.Execute()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "An internal server error occurred"})
	}

	return c.Status(http.StatusOK).JSON(events)
}

func (h *EventHandler) Create(c fiber.Ctx) error {
	var req create.Request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	validate, err := helper.Validate(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(validate)
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
		ID: event.ID,
	}

	return c.Status(http.StatusCreated).JSON(response)
}
