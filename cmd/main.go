package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"taxi-service/config"
	"taxi-service/internal/db"
	handler "taxi-service/internal/handlers"
	"taxi-service/internal/models"
	"taxi-service/internal/user"
)

func main() {

	cfg := config.LoadConfig()

	log.Println("Hello")

	// Initialize database
	db.InitDatabase(cfg)

	//// Crear repositorios
	userRepo := user.NewUserRepository(db.DB)
	//
	//// Crear servicios
	userService := user.NewUserService(userRepo)
	//
	//// Crear handlers
	userHandler := handler.NewUserHandler(userService)
	//
	router := gin.Default()
	//
	//// Rutas para Users
	router.GET("/users", userHandler.ListUsers)
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)

	// Migrate the schema
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	// Iniciar servidor
	log.Println("Iniciando servidor en :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}

}
