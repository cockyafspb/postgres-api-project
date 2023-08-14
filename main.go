package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"postgres-api-project/models"
	"postgres-api-project/repositories"
	"postgres-api-project/storage"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}
	err = models.MigrateOwners(db)

	if err != nil {
		log.Fatal("could not migrate db")
	}

	catR := repositories.CatRepository{
		DB: db,
	}

	ownerR := repositories.OwnerRepository{
		DB: db,
	}

	app := fiber.New()
	catR.SetupRoutes(app)
	ownerR.SetupRoutes(app)
	app.Listen(":8080")
}
