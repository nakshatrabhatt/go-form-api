// models/jwt_claim.go
package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaim extends the standard jwt.RegisteredClaims
type JWTClaim struct {
	jwt.RegisteredClaims
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
