// auth/jwt.go
package auth

import (
	"errors"	
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nakshatrabhatt/go-form-api/models"
)

// GenerateJWT generates a JWT token for a user
func GenerateJWT(userID uint, username, email string) (string, time.Time, error) {
	// JWT expiration time - 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)
	
	claims := &models.JWTClaim{
		UserID:   userID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Generate encoded token using the secret signing key
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	
	return tokenString, expirationTime, err
}

// Validating the JWT token
func ValidateToken(signedToken string) (*models.JWTClaim, error) {
	// Parse the JWT string and store the result in `claims`
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	
	if err != nil {
		return nil, err
	}
	
	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}
	
	// Check if the token is expired
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token is expired")
	}
	
	return claims, nil
}