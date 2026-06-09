package users

import (
	"database/sql"
	"log"

	util "github.com/mjl776/sports-management-platform/internal/utils"
)

type User struct {
	UserStatus string `json:"user_status"`
	UserID string `json:"user_id"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func NewUserObject(userStatus, email, password string) *User {

	userId := util.GenerateRandomULID()

	return &User {
		UserStatus: userStatus,
		UserID: userId.String(),
		Email: email,
		Password: password,
	}
}

func (s *UserService) CreateUsersTable() error {
	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_id VARCHAR(26) NOT NULL UNIQUE,
		user_status VARCHAR(30) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
        password_hash BYTEA NOT NULL
	)`

	_, err := s.db.Exec(createUserTableQuery)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err);
	}
	log.Println("Users table created successfully!")
	return nil
}

func (s *UserService) CreateUser(user User) error {
	insertUserQuery := `
	INSERT INTO users (user_status, user_id, password_hash, email)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	var userID string

	// Generate a secure password hash
	passwordHash, err := HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
	}

	// Insert the user into the database
	if err := s.db.QueryRow(insertUserQuery, user.UserStatus, user.UserID, passwordHash, user.Email).Scan(&userID); err != nil {
		log.Println("Error creating user:", err)
	}

	log.Println("User created successfully with ID:", userID)
	return nil
}
