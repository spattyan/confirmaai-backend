package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	application "github.com/spattyan/confirmaai-backend/cmd/app"
	"github.com/spattyan/confirmaai-backend/helper"
	eventDomain "github.com/spattyan/confirmaai-backend/internal/events/domain"
	participantDomain "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	environment, err := helper.SetupEnv()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: helper.HandleError,
	})

	database, err := gorm.Open(postgres.Open(environment.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to database")

	// migrations
	err = database.AutoMigrate(&eventDomain.Event{}, &eventDomain.EventRole{}, &userDomain.User{}, &participantDomain.Participant{}, &participantDomain.Permission{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	auth := helper.SetupAuth(environment.AuthToken)

	application.SetupApplication(database, auth, app)

	if err := app.Listen(environment.ServerPort); err != nil {
		log.Printf("Error starting server: %s", err)
	}

}
