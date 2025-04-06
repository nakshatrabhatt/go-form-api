package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nakshatrabhatt/go-form-api/controllers"
	"github.com/nakshatrabhatt/go-form-api/database"
	"github.com/nakshatrabhatt/go-form-api/middleware"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Set a default JWT secret if not provided in environment
	if os.Getenv("JWT_SECRET") == "" {
		log.Println("Warning: JWT_SECRET not set, using default (insecure)")
		os.Setenv("JWT_SECRET", "default_jwt_secret_change_this_in_production")
	}

	// Initialize database connection
	database.ConnectDB()

	// Set up Gin router
	router := gin.Default()

	// CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// User routes
		protected.GET("/profile", controllers.GetUserProfile)

		// Form routes - adapted to match your original forms.go approach
		protected.POST("/forms", controllers.CreateForm)
		protected.GET("/forms", controllers.GetForms)        // Query parameter version
		protected.GET("/forms/:id", controllers.GetFormByID) // Path parameter version
		protected.PUT("/forms/:id", controllers.UpdateForm)
		protected.DELETE("/forms/:id", controllers.DeleteForm)
	}

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	log.Printf("Server running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
