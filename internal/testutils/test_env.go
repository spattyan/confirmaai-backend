package testutils

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"

	eventDomain "github.com/spattyan/confirmaai-backend/internal/events/domain"
	participantDomain "github.com/spattyan/confirmaai-backend/internal/participants/domain"
	userDomain "github.com/spattyan/confirmaai-backend/internal/users/domain"

	"github.com/testcontainers/testcontainers-go"
	pgContainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	once      sync.Once
	container testcontainers.Container
)

func SetupTestEnv() {
	once.Do(func() {
		ctx := context.Background()

		postgresContainer, err := pgContainer.Run(ctx,
			"postgres:latest",
			pgContainer.WithDatabase("testdb"),
			pgContainer.WithUsername("postgres"),
			pgContainer.WithPassword("secret"),
			pgContainer.BasicWaitStrategies(),
		)

		if err != nil {
			log.Fatalf("❌ failed to start postgres container: %v", err)
		}

		host, _ := postgresContainer.Host(ctx)
		port, _ := postgresContainer.MappedPort(ctx, "5432")

		dsn := fmt.Sprintf("host=%s user=postgres password=secret dbname=testdb port=%s sslmode=disable", host, port.Port())

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ failed to connect to postgres: %v", err)
		}

		if err := DB.AutoMigrate(&eventDomain.Event{}, &userDomain.User{}, &participantDomain.Participant{}, &eventDomain.EventRole{}); err != nil {
			log.Fatalf("❌ failed to migrate: %v", err)
		}

		log.Println("✅ Postgres container ready and migrations done")
	})
}

func ResetDB() {
	if DB == nil {
		log.Fatalf("❌ DB not initialized. Did you forget to call SetupTestEnv()?")
	}

	tables := []string{
		"participants",
		"events",
		"users",
		"event_roles",
	}

	for _, table := range tables {
		if err := DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE;", table)).Error; err != nil {
			log.Fatalf("❌ failed to truncate table %s: %v", table, err)
		}
	}

	log.Println("🧹 Database cleaned successfully")
}

func TeardownTestEnv() {
	if container != nil {
		_ = container.Terminate(context.Background())
	}
}

func SetupIsolatedTest(t *testing.T) {
	SetupTestEnv()
	ResetDB()
	t.Cleanup(func() {
		ResetDB()
	})
}
