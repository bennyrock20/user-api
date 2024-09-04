package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"taxi-service/config"
	"taxi-service/internal/db"
	handler "taxi-service/internal/handlers"
	"taxi-service/internal/models"
	"taxi-service/internal/user"
	"taxi-service/middlewares"
)

func main() {

	cfg := config.LoadConfig()

	log.Println("Hello")

	// Initialize database
	db.InitDatabase(cfg)

	userRepo := user.NewUserRepository(db.DB)
	userService := user.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Routers
	public := gin.Default()

	// Public route
	public.POST("/login", handler.LoginHandler)

	// Protected routes
	protected := public.Group("/api")
	protected.Use(middlewares.AuthMiddleware())

	log.Printf("CHekcing routes")

	////  Users Routes
	protected.GET("/users", userHandler.ListUsers)
	protected.GET("/users/:id", userHandler.GetUser)
	//protected.POST("/users", userHandler.CreateUser)
	protected.PUT("/users/:id", userHandler.UpdateUser)
	protected.DELETE("/users/:id", userHandler.DeleteUser)

	// Migrate the schema
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Iniciando servidor en :8080")
	if err := public.Run(":8080"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}

}
