package main

import (
	_ "hostflow/profile-service/docs" // Import generated swagger docs
	"hostflow/profile-service/internal/bootstrap"
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// @title Hostflow Profile Service API
// @version 1.0
// @description This is a comprehensive profile service API for managing users and organizations.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@hostflow.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http https

// @tag.name users
// @tag.description Operations related to users

func main() {
	if err := godotenv.Load("./dev.env"); err != nil {
		log.Printf("Warning: Error loading dev.env file: %v", err)
	}

	fx.New(
		bootstrap.Module,
	).Run()
}
