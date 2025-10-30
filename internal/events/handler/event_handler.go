package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createevent"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createeventrole"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/geteventbyid"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/geteventroles"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/listevents"
)

type EventHandler struct {
	auth                *helper.Auth
	createEventUseCase  createevent.UseCase
	listEventsUseCase   listevents.UseCase
	getEventByIdUseCase geteventbyid.UseCase
	createEventRole     createeventrole.UseCase
	getEventRoles       geteventroles.UseCase
}

func NewEventHandler(
	auth *helper.Auth,
	createEventUC createevent.UseCase,
	listEventsUC listevents.UseCase,
	getEventByIdUC geteventbyid.UseCase,
	createEventRoleUC createeventrole.UseCase,
	getEventRolesUC geteventroles.UseCase) *EventHandler {

	return &EventHandler{
		auth:                auth,
		createEventUseCase:  createEventUC,
		listEventsUseCase:   listEventsUC,
		getEventByIdUseCase: getEventByIdUC,
		createEventRole:     createEventRoleUC,
		getEventRoles:       getEventRolesUC,
	}
}

func (h *EventHandler) EventRoutes(router fiber.Router) {
	eventRoutes := router.Group("/events", h.auth.Authorize)
	eventRoutes.Post("/new", h.CreateEvent)
	eventRoutes.Get("/", h.ListEvents)
	eventRoutes.Get("/:id", h.GetEventById)

	eventRoutes.Post("/roles/new", h.CreateEventRole)
	eventRoutes.Get("/roles/:id", h.ListEventRoles)
}

func (h *EventHandler) GetEventById(c fiber.Ctx) error {
	req, err := helper.BindAndValidate[geteventbyid.DTO](c, helper.BindURI)

	if err != nil {
		return err
	}

	events, err := h.getEventByIdUseCase.Execute(geteventbyid.DTO{
		Id: req.Id,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(events)
}

func (h *EventHandler) ListEvents(c fiber.Ctx) error {
	_, err := helper.BindAndValidate[listevents.DTO](c, helper.BindQuery)

	if err != nil {
		return err
	}

	events, err := h.listEventsUseCase.Execute(listevents.DTO{})

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(events)
}

func (h *EventHandler) CreateEvent(c fiber.Ctx) error {

	req, err := helper.BindAndValidate[createevent.DTO](c, helper.BindBody)

	if err != nil {
		return err
	}

	user := h.auth.GetCurrentUser(c)

	if user.ID == uuid.Nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	event, err := h.createEventUseCase.Execute(createevent.DTO{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		DateAndTime:      req.DateAndTime,
		ParticipantLimit: req.ParticipantLimit,
		User:             &user,
	})

	if err != nil {
		return err
	}

	response := createevent.Response{
		ID: event.ID,
	}

	return c.Status(http.StatusCreated).JSON(response)
}

func (h *EventHandler) CreateEventRole(c fiber.Ctx) error {
	req, err := helper.BindAndValidate[createeventrole.DTO](c, helper.BindBody)

	if err != nil {
		return err
	}
	user := h.auth.GetCurrentUser(c)

	if user.ID == uuid.Nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	eventRole, err := h.createEventRole.Execute(createeventrole.DTO{
		EventID: req.EventID,
		Name:    req.Name,
		Slots:   req.Slots,
		User:    &user,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"id": eventRole.ID,
	})

}

func (h *EventHandler) ListEventRoles(c fiber.Ctx) error {
	req, err := helper.BindAndValidate[geteventroles.DTO](c, helper.BindURI)

	if err != nil {
		return err
	}

	roles, err := h.getEventRoles.Execute(geteventroles.DTO{
		Id: req.Id,
	})

	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(roles)
}
