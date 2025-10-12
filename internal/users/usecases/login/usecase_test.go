package login

import (
	"testing"

	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/testutils"
	"github.com/spattyan/confirmaai-backend/internal/users/repository"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/register"
)

func TestLoginUseCase(t *testing.T) {
	testutils.SetupIsolatedTest(t)
	auth := helper.SetupAuth("testing")

	repo := repository.NewGormRepository(testutils.DB)
	regUC := register.NewUseCase(repo, auth)
	_, err := regUC.Execute(register.DTO{
		Email:    "testing@email.com",
		Name:     "Login User",
		Password: "abc123",
	})
	if err != nil {
		t.Fatalf("❌ failed to create user for login test: %v", err)
	}

	t.Run("login a user", func(t *testing.T) {
		repo := repository.NewGormRepository(testutils.DB)
		usecase := NewUseCase(repo, helper.SetupAuth("testing"))

		t.Run("with valid credentials", func(t *testing.T) {
			result, err := usecase.Execute(DTO{
				Email:    "testing@email.com",
				Password: "abc123",
			})

			if err != nil {
				t.Errorf("❌ expected no error, got %v", err)
				return
			}

			t.Logf("✅ Successfully logged user with ID and token: %v, %v", result.ID, result.Token)
		})

		t.Run("with invalid credentials", func(t *testing.T) {
			_, err := usecase.Execute(DTO{
				Email:    "testing@email.com",
				Password: "abc1232",
			})

			if err == nil {
				t.Errorf("❌ expected an error, but got none")
				return
			}

			t.Logf("✅ Login denied because credentials are invalid, err: %v", err.Error())
		})
	})
}
