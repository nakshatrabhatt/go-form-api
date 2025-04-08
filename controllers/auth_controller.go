package controllers

import (
	"fmt"
	"io"
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nakshatrabhatt/go-form-api/auth"
	"github.com/nakshatrabhatt/go-form-api/database"
	"github.com/nakshatrabhatt/go-form-api/models"
	"github.com/nakshatrabhatt/go-form-api/utils"

	"log"
)

// Login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Register request body
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Response after successful authentication
type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
	UserID    uint   `json:"userId"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

// User registration
func Register(c *gin.Context) {
	// Debug request headers
	log.Println("Request Headers:")
	for name, values := range c.Request.Header {
		log.Printf("%s: %v", name, values)
	}

	// Read the request body but keep it for later binding
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
		return
	}
	
	// Log the raw request body for debugging
	bodyString := string(bodyBytes)
	log.Printf("Raw request body: %s", bodyString)
	
	// Restore the request body for binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var request RegisterRequest
	// Validate request body
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error binding JSON: %v", err)})
		return
	}

	log.Printf("Received registration request: username=%s, email=%s", request.Username, request.Email)

	// Check if user with the same email exists
	var existingUser models.User
	if database.DB.Where("email = ?", request.Email).First(&existingUser).RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	// Check if username is already taken
	if database.DB.Where("username = ?", request.Username).First(&existingUser).RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Username is already taken"})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		log.Println("Error saving user to DB:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Generate JWT token
	token, expiresAt, err := auth.GenerateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return response
	c.JSON(http.StatusCreated, AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt.Format(http.TimeFormat),
		UserID:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
	})
}

// Login handles user login
func Login(c *gin.Context) {
	var request LoginRequest

	// Validate request body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if database.DB.Where("email = ?", request.Email).First(&user).RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if !utils.CheckPasswordHash(request.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, expiresAt, err := auth.GenerateJWT(user.ID, user.Username, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt.Format(http.TimeFormat),
		UserID:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
	})
}

// Returns the profile of the authenticated user
func GetUserProfile(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Find user by ID
	var user models.User
	if database.DB.First(&user, userID).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return user profile (password is already hidden by struct tag)
	c.JSON(http.StatusOK, user)
}