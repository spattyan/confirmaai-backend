package users_test

import (
	"testing"

	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/testutils"
	"github.com/spattyan/confirmaai-backend/internal/users/repository"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/login"
	"github.com/spattyan/confirmaai-backend/internal/users/usecases/register"
)

func TestUserFlow(t *testing.T) {
	testutils.SetupTestEnv()

	t.Run("register a user", func(t *testing.T) {
		repo := repository.NewGormRepository(testutils.DB)
		usecase := register.NewUseCase(repo, helper.SetupAuth("testing"))

		t.Run("with valid fields", func(t *testing.T) {
			result, err := usecase.Execute(register.DTO{
				Email:    "testing@email.com",
				Name:     "Testing User",
				Password: "abc123",
			})

			if err != nil {
				t.Errorf("❌ expected no error, got %v", err)
				return
			}

			t.Logf("✅ Successfully created user with ID and token: %v, %v", result.ID, result.Token)
		})

		t.Run("with same email", func(t *testing.T) {
			_, err := usecase.Execute(register.DTO{
				Email:    "testing@email.com",
				Name:     "Testing User 2",
				Password: "abc123123",
			})

			if err == nil {
				t.Errorf("❌ expected an error, but got none")
				return
			}

			t.Logf("✅ User creation failed with same email, err: %v", err.Error())
		})

		t.Run("with invalid fields", func(t *testing.T) {
			_, err := usecase.Execute(register.DTO{
				Email:    "testing@emailcom",
				Name:     "Testing@&#ˆ*E(User",
				Password: "ab",
			})

			if err == nil {
				t.Errorf("❌ expected an error, but got none")
				return
			}

			t.Logf("✅ User creation failed with invalid fields, err: %v", err.Error())
		})

	})

	t.Run("login a user", func(t *testing.T) {
		repo := repository.NewGormRepository(testutils.DB)
		usecase := login.NewUseCase(repo, helper.SetupAuth("testing"))

		t.Run("with valid credentials", func(t *testing.T) {
			result, err := usecase.Execute(login.DTO{
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
			_, err := usecase.Execute(login.DTO{
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
