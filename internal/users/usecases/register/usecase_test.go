package register

import (
	"testing"

	"github.com/spattyan/confirmaai-backend/helper"
	"github.com/spattyan/confirmaai-backend/internal/testutils"
	"github.com/spattyan/confirmaai-backend/internal/users/repository"
)

func TestRegisterUseCase(t *testing.T) {
	testutils.SetupIsolatedTest(t)

	t.Run("register a user", func(t *testing.T) {
		repo := repository.NewGormRepository(testutils.DB)
		usecase := NewUseCase(repo, helper.SetupAuth("testing"))

		t.Run("with valid fields", func(t *testing.T) {
			result, err := usecase.Execute(DTO{
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
			_, err := usecase.Execute(DTO{
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
			_, err := usecase.Execute(DTO{
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
}
