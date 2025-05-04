package users

import (
	"database/sql"
	"log"
)

type User struct {
	UserStatus string `json:"user_status"`
	EmployeeId string `json:"employee_id"`
	Password string `json:"password_hash"`
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func NewUserObject(userStatus, employeeid, password string) *User {
	return &User {
		UserStatus: userStatus,
		EmployeeId: employeeid,
		Password: password,
	}
}

func (s *UserService) CreateUsersTable() error {
	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_status VARCHAR(30) NOT NULL,
		employee_id VARCHAR(30) NOT NULL,
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
	INSERT INTO users (user_status, employee_id, password_hash)
	VALUES ($1, $2, $3)
	RETURNING uid;
	`
	var userID string
	log.Println("Creating user with status:", user.UserStatus, "and employee ID:", user.EmployeeId)
	// Generate a secure password hash
	passwordHash, err := HashPassword(user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
	}

	// Insert the user into the database
	if err := s.db.QueryRow(insertUserQuery, user.UserStatus, user.EmployeeId, passwordHash).Scan(&userID); err != nil {
		log.Println("Error creating user:", err)
	}

	log.Println("User created successfully with ID:", userID)
	return nil
}
