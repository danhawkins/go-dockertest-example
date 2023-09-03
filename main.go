package main

import (
	"log"

	"github.com/danhawkins/go-dockertest-example/database"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.Connect()

	database.CreatePerson()

	count := database.CountPeople()

	log.Printf("Database has %d people", count)
}
