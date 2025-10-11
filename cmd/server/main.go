package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	eventDomain "github.com/spattyan/confirmaai-backend/internal/events/domain"
	eventHand "github.com/spattyan/confirmaai-backend/internal/events/handler"
	eventRepo "github.com/spattyan/confirmaai-backend/internal/events/repository"
	participantDomain "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/repository"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
	userHand "github.com/spattyan/confirmaai-backend/internal/users/handler"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	environment, err := helper.SetupEnv()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println(environment)
	app := fiber.New()

	database, err := gorm.Open(postgres.Open(environment.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to database")

	// migrations
	err = database.AutoMigrate(&eventDomain.Event{}, &eventDomain.EventRole{}, &userDomain.User{}, &participantDomain.Participant{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	auth := helper.SetupAuth(environment.AuthToken)

	eventRepository := eventRepo.NewGormRepository(database)
	userRepository := userRepo.NewGormRepository(database)
	participantRepository := participantRepo.NewGormRepository(database)

	fmt.Println(participantRepository) // to avoid unused variable error

	eventHandler := eventHand.NewEventHandler(eventRepository, userRepository, participantRepository, auth)
	eventHandler.EventRoutes(app)

	userHandler := userHand.NewUserHandler(userRepository, auth)
	userHandler.UserRoutes(app)

	if err := app.Listen(environment.ServerPort); err != nil {
		log.Printf("Error starting server: %s", err)
	}

}
