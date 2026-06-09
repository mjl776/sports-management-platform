package users

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secret_key = []byte(os.Getenv("JWT_SECRET_KEY"))

// HashPassword hashes the password using bcrypt
func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}

func (s *UserService) VerifyPasswordAndUserStatus(email, plainPassword string) (string, string, error) {
    email = strings.TrimSpace(strings.ToLower(email))
    plainPassword = strings.TrimSpace(plainPassword)

    var hashedPassword []byte
    var userStatus string
    var userID string
	fmt.Printf("Verifying user with email: %s\n", email)
    query := `SELECT user_id, password_hash, user_status FROM users WHERE email = $1`
    err := s.db.QueryRow(query, email).Scan(&userID, &hashedPassword, &userStatus)
    if err != nil {
        return "", "", fmt.Errorf("failed to retrieve user: %w", err)
    }

    if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword)); err != nil {
        return "", "", fmt.Errorf("invalid password")
    }

    return userID, userStatus, nil
}


func (s *UserService) AuthenticationLogin(email, password string) (string, error) {
	var err error
	// check if password authentication works exist
	userStatus, userId, err := s.VerifyPasswordAndUserStatus(email, password)
	if err != nil {
		return "", fmt.Errorf("failed to verify password: %w", err)
	}

	// Generate JWT token
	token, err := s.GenerateJWT(userId, userStatus)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %w", err)
	}

	// return token
	return token, nil

}


func (s *UserService) GenerateJWT(userId, userStatus string) (string, error) {
	// Create a new JWT token
	claims := jwt.MapClaims{
        "user_id": userId,
        "role":        userStatus,
        "exp":         time.Now().Add(time.Hour * 24).Unix(),
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret_key)
}


