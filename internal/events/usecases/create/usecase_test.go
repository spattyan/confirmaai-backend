package create

import (
	"testing"

	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/events/repository"
	"github.com/spattyan/confirmaai-backend/internal/events/usecases/createEventRole"
	participantRepo "github.com/spattyan/confirmaai-backend/internal/participants/repository"
	"github.com/spattyan/confirmaai-backend/internal/participants/usecases/createParticipant"
	"github.com/spattyan/confirmaai-backend/internal/testutils"
	userRepo "github.com/spattyan/confirmaai-backend/internal/users/repository"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/register"
)

func TestCreateUseCase(t *testing.T) {
	testutils.SetupIsolatedTest(t)
	auth := helper.SetupAuth("testing")

	repo := userRepo.NewGormRepository(testutils.DB)
	regUC := register.NewUseCase(repo, auth)
	userUs, err := regUC.Execute(register.DTO{
		Email:    "testing@email.com",
		Name:     "Login User",
		Password: "abc123",
	})
	if err != nil {
		t.Fatalf("❌ failed to create user for event test: %v", err)
	}

	t.Run("create a event", func(t *testing.T) {
		repo := repository.NewGormRepository(testutils.DB)
		userRepository := userRepo.NewGormRepository(testutils.DB)
		participantRepository := participantRepo.NewGormRepository(testutils.DB)

		createEventRoleUs := createEventRole.NewUseCase(repo)
		createParticipantUs := createParticipant.NewUseCase(userRepository, repo, participantRepository)

		usecase := NewUseCase(repo, createEventRoleUs, createParticipantUs)

		user, err := userRepository.FindByID(userUs.ID)

		if err != nil {
			t.Fatalf("❌ failed to find user by ID for event test: %v", err)
		}

		t.Run("with valid fields", func(t *testing.T) {
			result, err := usecase.Execute(DTO{
				Title:            "Testing Event Title",
				Description:      "Testing Event Description",
				Location:         "Testing Event Location",
				DateAndTime:      "2026-02-02 16:40:00",
				ParticipantLimit: 5,
				User:             user,
			})

			if err != nil {
				t.Errorf("❌ expected no error, got %v", err)
				return
			}

			t.Logf("✅ Successfully created event, ID: %v", result.ID)
		})

		t.Run("with invalid date", func(t *testing.T) {
			_, err := usecase.Execute(DTO{
				Title:            "Testing Event Title",
				Description:      "Testing Event Description",
				Location:         "Testing Event Location",
				DateAndTime:      "2026-02-02 16:40",
				ParticipantLimit: 5,
				User:             user,
			})

			if err == nil {
				t.Errorf("❌ expected error, got no one")
				return
			}

			t.Logf("✅ Successfully denied event creation, err: %v", err.Error())
		})

		t.Run("with invalid participant limit", func(t *testing.T) {
			_, err := usecase.Execute(DTO{
				Title:            "Testing Event Title",
				Description:      "Testing Event Description",
				Location:         "Testing Event Location",
				DateAndTime:      "2026-02-02 16:40:00",
				ParticipantLimit: 0,
				User:             user,
			})

			if err == nil {
				t.Errorf("❌ expected error, got no one")
				return
			}

			t.Logf("✅ Successfully denied event creation, err: %v", err.Error())
		})
	})
}
