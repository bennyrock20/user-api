package cmd

import (
	"log"
	"net/http"
	"taxi-service/config"
	"taxi-service/internal/db"

	"taxi-service/internal/models"
	"taxi-service/middlewares"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func run_server() {

	cfg := config.LoadConfig()

	log.Println("Hello")

	// Initialize database
	db.InitDatabase(cfg)

	// userRepo := user.NewUserRepository(db.DB)
	// userService := user.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userService)

	// Routers
	public := gin.Default()

	// Public route
	public.GET("/about/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hellow World"})
	})

	// Protected routes
	protected := public.Group("/api/v1")

	// Apply middleware
	protected.Use(middlewares.JWTAuthMiddleware())

	protected.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(http.StatusOK, gin.H{"message": "Welcome", "userID": userID})
	})

	// Migrate the schema
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Iniciando servidor en :8080")
	if err := public.Run(":8080"); err != nil {
		log.Fatalf("No se pudo iniciar el servidor: %v", err)
	}

}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the web server using Gin",
	Long:  `This command starts the Gin web server to handle HTTP requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the Gin web server...")

		run_server()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
