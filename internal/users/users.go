package users

import (
	"database/sql"
	"log"
)

type User struct {
	ID          string `json:"id"`
	User_status string `json:"user_status"`
	Employee_id string `json:"employee_id"`
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func NewUserObject(id, userStatus, employeeid string) *User {
	return &User {
		ID:          id,
		User_status: userStatus,
		Employee_id: employeeid,
	}
}

func (s *UserService) CreateUsersTable() error {
	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_status VARCHAR(30) NOT NULL,
		employee_id VARCHAR(30) NOT NULL
	)`

	_, err := s.db.Exec(createUserTableQuery)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err);
	}
	log.Println("Users table created successfully!")
	return nil
}

