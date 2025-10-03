package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
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

	app := fiber.New()

	database, err := gorm.Open(postgres.Open("host=127.0.0.1 user=root password=root dbname=confirmaai port=5432 sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	log.Println("Successfully connected to database")

	// migrations
	err = database.AutoMigrate(&eventDomain.Event{}, &userDomain.User{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}

	eventRepository := eventRepo.NewGormRepository(database)
	eventHandler := eventHand.NewEventHandler(eventRepository)

	eventHandler.EventRoutes(app)

	userRepository := userRepo.NewGormRepository(database)
	userHandler := userHand.NewUserHandler(userRepository)

	userHandler.UserRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Printf("Error starting server: %s", err)
	}

}
