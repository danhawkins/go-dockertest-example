package database_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/danhawkins/go-dockertest-example/database"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	// Start a new docker pool
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// Uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	pg, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "15",
		Env: []string{
			"POSTGRES_DB=example",
			"POSTGRES_HOST_AUTH_METHOD=trust",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	pg.Expire(10)

	// Set this so our app can use it
	postgresPort := pg.GetPort("5432/tcp")
	os.Setenv("POSTGRES_PORT", postgresPort)

	// Wait for the HTTP endpoint to be ready
	if err := pool.Retry(func() error {
		_, connErr := gorm.Open(postgres.Open(fmt.Sprintf("postgresql://postgres@localhost:%s/example", postgresPort)), &gorm.Config{})
		if connErr != nil {
			return connErr
		}

		return nil
	}); err != nil {
		panic("Could not connect to postgres: " + err.Error())
	}

	code := m.Run()

	os.Exit(code)
}

func TestCreatePerson(t *testing.T) {
	// Connect to the database
	database.Connect()

	// Create a person in the database
	database.CreatePerson()

	// Check that the person was created
	count := database.CountPeople()

	if count != 1 {
		t.Errorf("Expected 1 person to be in the database, got %d", count)
	}
}
