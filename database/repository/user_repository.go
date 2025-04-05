package repository

import (
    "github.com/nakshatrabhatt/go-form-api/database"
    "github.com/nakshatrabhatt/go-form-api/models"
)

// GetUserByEmail finds a user by email address
func GetUserByEmail(email string) (models.User, error) {
    var user models.User
    result := database.DB.Where("email = ?", email).First(&user)
    return user, result.Error
}

// GetUserByID finds a user by ID
func GetUserByID(id uint) (models.User, error) {
    var user models.User
    result := database.DB.First(&user, id)
    return user, result.Error
}

// CreateUser creates a new user in the database
func CreateUser(user models.User) (models.User, error) {
    result := database.DB.Create(&user)
    return user, result.Error
}

// UpdateUser updates an existing user
func UpdateUser(user models.User) error {
    result := database.DB.Save(&user)
    return result.Error
}