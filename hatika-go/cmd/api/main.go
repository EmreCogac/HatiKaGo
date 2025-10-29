package main

import (
	"fmt"
	"log"

	"hatika-go/internal/application/services"
	"hatika-go/internal/infrastructure/config"
	"hatika-go/internal/infrastructure/persistence"
	"hatika-go/internal/interfaces/http"
	"hatika-go/internal/interfaces/http/handlers"

	_ "hatika-go/docs"
)

func main() {
	log.Println("Starting LLMOCR API Server...")

	cfg, err := config.LoadConfig("./config")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Printf("Configuration loaded successfully")

	dbConfig := &persistence.DatabaseConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
		SSLMode:  cfg.Database.SSLMode,
	}

	db, err := persistence.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")

	if err := persistence.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	if err := persistence.SeedData(db); err != nil {
		log.Printf("Warning: Failed to seed data: %v", err)
	}

	projectRepo := persistence.NewProjectRepository(db)

	// Initialize services
	projectService := services.NewProjectService(projectRepo)

	// Initialize handlers
	projectHandler := handlers.NewProjectHandler(projectService)

	// Setup router
	router := http.SetupRouter(projectHandler)

	// Start server
	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", address)
	log.Printf("Swagger documentation: http://localhost:%d/swagger/index.html", cfg.Server.Port)

	if err := router.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
