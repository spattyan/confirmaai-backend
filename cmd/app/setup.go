package application

import (
	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	eventHand "github.com/spattyan/confirmaai-backend/internal/events/handler"
	userHand "github.com/spattyan/confirmaai-backend/internal/users/handler"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/login"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/register"

	eventRepo "github.com/spattyan/confirmaai-backend/internal/events/repository"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createevent"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createeventrole"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/geteventbyid"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/geteventroles"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/listevents"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/repository"
	"github.com/spattyan/confirmaai-backend/internal/participants/usecases/createparticipant"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/repository"
	"gorm.io/gorm"
)

func SetupApplication(database *gorm.DB, auth *helper.Auth, app *fiber.App) {
	// === REPOSITORIES ===
	eventRepository := eventRepo.NewGormRepository(database)
	userRepository := userRepo.NewGormRepository(database)
	participantRepository := participantRepo.NewGormRepository(database)

	// === USE CASES ===
	listEventsUseCase := listevents.NewUseCase(eventRepository)
	getEventByIdUseCase := geteventbyid.NewUseCase(eventRepository)
	getEventRolesUseCase := geteventroles.NewUseCase(eventRepository)
	createEventRoleUseCase := createeventrole.NewUseCase(eventRepository)
	createParticipantUseCase := createparticipant.NewUseCase(userRepository, eventRepository, participantRepository)

	registerUseCase := register.NewUseCase(userRepository, auth)
	loginUseCase := login.NewUseCase(userRepository, auth)

	// === ORCHESTRATORS USE CASES ===
	createEventUseCase := createevent.NewUseCase(eventRepository, createEventRoleUseCase, createParticipantUseCase)

	// === HANDLERS ===
	eventHandler := eventHand.NewEventHandler(
		auth,
		createEventUseCase,
		listEventsUseCase,
		getEventByIdUseCase,
		createEventRoleUseCase,
		getEventRolesUseCase,
	)
	eventHandler.EventRoutes(app)

	userHandler := userHand.NewUserHandler(registerUseCase, loginUseCase)
	userHandler.UserRoutes(app)
}
