package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/spattyan/confirmaai-backend/helper"
	eventDomain "github.com/spattyan/confirmaai-backend/internal/events/domain"
	eventHand "github.com/spattyan/confirmaai-backend/internal/events/handler"
	eventRepo "github.com/spattyan/confirmaai-backend/internal/events/repository"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"
	userHand "github.com/spattyan/confirmaai-backend/internal/users/handler"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello World")

	enviroment, err := helper.SetupEnv()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println(enviroment)
	app := fiber.New()

	database, err := gorm.Open(postgres.Open(enviroment.Dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to database")

	// migrations
	err = database.AutoMigrate(&eventDomain.Event{}, &userDomain.User{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	auth := helper.SetupAuth(enviroment.AuthToken)

	eventRepository := eventRepo.NewGormRepository(database)
	eventHandler := eventHand.NewEventHandler(eventRepository, auth)

	eventHandler.EventRoutes(app)

	userRepository := userRepo.NewGormRepository(database)
	userHandler := userHand.NewUserHandler(userRepository, auth)

	userHandler.UserRoutes(app)

	if err := app.Listen(enviroment.ServerPort); err != nil {
		log.Printf("Error starting server: %s", err)
	}

}
