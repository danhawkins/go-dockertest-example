package database_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/danhawkins/go-dockertest-example/database"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
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

	postgres, err := pool.RunWithOptions(&dockertest.RunOptions{
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

	postgres.Expire(10)

	// Set this so our app can use it
	postgresPort := postgres.GetPort("5432/tcp")
	os.Setenv("POSTGRES_PORT", postgresPort)

	// Wait 1 second for the container to start
	time.Sleep(1 * time.Second)

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
