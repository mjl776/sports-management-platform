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

func (s *UserService) VerifyPasswordAndUserStatus(employeeID, plainPassword string) (string, error) {
    var hashedPassword []byte

	// Trim any leading or trailing whitespace from the employee ID and password
	plainPassword = strings.TrimSpace(plainPassword)

    // Retrieve the hashed password from the database
    passQuery := `SELECT password_hash FROM users WHERE employee_id = $1`
    if err := s.db.QueryRow(passQuery, employeeID).Scan(&hashedPassword); err != nil {
        return "", fmt.Errorf("failed to retrieve password hash: %w", err)
    }

    // Compare the hashed password with the plain password
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(plainPassword)); err != nil && hashedPassword != nil {
        return "", fmt.Errorf("invalid password: %w", err)
    }

	var userStatus string
	// Retrieve the user status from the database
	userStatusQuery := `SELECT user_status FROM users WHERE employee_id = $1`
	if 	err := s.db.QueryRow(userStatusQuery, employeeID).Scan(&userStatus); err != nil{
		return "", fmt.Errorf("failed to retrieve user status: %w", err)
	}
	return userStatus, nil

}


func (s *UserService) AuthenticationLogin(employeeID, password string) (string, error) {
	var err error
	// check if password authentication works exist
	userStatus, err := s.VerifyPasswordAndUserStatus(employeeID, password)
	if err != nil {
		return "", fmt.Errorf("failed to verify password: %w", err)
	}

	// Generate JWT token
	token, err := s.GenerateJWT(employeeID, userStatus)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %w", err)
	}

	// return token
	return token, nil

}


func (s *UserService) GenerateJWT(employeeID, role string) (string, error) {
	// Create a new JWT token
	claims := jwt.MapClaims{
        "employee_id": employeeID,
        "role":        role,
        "exp":         time.Now().Add(time.Hour * 24).Unix(),
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret_key)
}


