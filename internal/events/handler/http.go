package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/create"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/getById"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/list"
)

type EventHandler struct {
	auth           *helper.Auth
	createUseCase  create.UseCase
	listUseCase    list.UseCase
	getByIdUseCase getById.UseCase
}

func NewEventHandler(repository domain.Repository, auth *helper.Auth) *EventHandler {
	return &EventHandler{auth: auth, createUseCase: create.NewUseCase(repository), listUseCase: list.NewUseCase(repository), getByIdUseCase: getById.NewUseCase(repository)}
}

func (h *EventHandler) EventRoutes(router fiber.Router) {
	eventRoutes := router.Group("/events", h.auth.Authorize)
	eventRoutes.Post("/new", h.Create)
	eventRoutes.Get("/", h.List)
	eventRoutes.Get("/:id", h.GetById)
}

func (h *EventHandler) GetById(c fiber.Ctx) error {
	var req getById.Request
	if err := c.Bind().URI(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request body"})
	}

	validate, err := helper.Validate(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(validate)
	}

	events, err := h.getByIdUseCase.Execute(getById.DTO{
		Id: req.Id,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "An internal server error occurred"})
	}

	return c.Status(http.StatusOK).JSON(events)
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

	//user := h.auth.GetCurrentUser(c)

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
