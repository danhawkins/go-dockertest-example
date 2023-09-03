package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Person struct {
		gorm.Model
		Name string
		Age  uint
	}
)

var db *gorm.DB

func Connect() {
	log.Println("Setting up the database")

	pgUrl := fmt.Sprintf("postgresql://postgres@127.0.0.1:%s/example", os.Getenv("POSTGRES_PORT"))
	log.Printf("Connecting to %s\n", pgUrl)
	var err error

	db, err = gorm.Open(postgres.Open(pgUrl), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})
}

func CreatePerson() {
	log.Println("Creating a new person in the database")
	person := Person{Name: "Danny", Age: 42}
	db.Create(&person)

	log.Println("Trying to write a new person to the database")
}

func CountPeople() int {
	var count int64
	db.Model(&Person{}).Count(&count)
	return int(count)
}
