package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

type (
	Person struct {
		gorm.Model
		Name string
		Age  uint
	}
)

func main() {
	log.Println("Setting up the database")

	pgUrl := fmt.Sprintf("postgresql://postgres@localhost:%s/example", os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(pgUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})

	log.Println("Creating a new person in the database")
	person := Person{Name: "Danny", Age: 42}
	db.Create(&person)

	log.Println("Trying to write a new person to the database")
}
