package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/domain"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/create"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createEventRole"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/getById"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/list"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	"github.com/spattyan/confirmaai-backend/internal/participants/usecases/createParticipant"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/domain"
)

type EventHandler struct {
	auth              *helper.Auth
	createUseCase     create.UseCase
	listUseCase       list.UseCase
	getByIdUseCase    getById.UseCase
	createEventRole   createEventRole.UseCase
	createParticipant createParticipant.UseCase
}

func NewEventHandler(repository domain.Repository, userRepository userRepo.Repository, participantRepo participantRepo.Repository, auth *helper.Auth) *EventHandler {
	createParticipantUs := createParticipant.NewUseCase(userRepository, repository, participantRepo)
	createEventRoleUs := createEventRole.NewUseCase(repository)

	return &EventHandler{
		auth:              auth,
		createUseCase:     create.NewUseCase(repository, createEventRoleUs, createParticipantUs),
		listUseCase:       list.NewUseCase(repository),
		getByIdUseCase:    getById.NewUseCase(repository),
		createEventRole:   createEventRoleUs,
		createParticipant: createParticipantUs,
	}
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

	user := h.auth.GetCurrentUser(c)

	if user.ID == uuid.Nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	event, err := h.createUseCase.Execute(create.DTO{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		DateAndTime:      req.DateAndTime,
		ParticipantLimit: req.ParticipantLimit,
		User:             &user,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": fmt.Sprintf("An internal server error occurred while creating event, %v", err)})
	}

	response := create.Response{
		ID: event.ID,
	}

	return c.Status(http.StatusCreated).JSON(response)
}
