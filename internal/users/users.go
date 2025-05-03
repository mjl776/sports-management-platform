package users

import (
	"database/sql"
	"log"
)

type User struct {
	ID          string `json:"id"`
	User_status string `json:"user_status"`
	Employee_id string `json:"employee_id"`
	Password_hash string `json:"password_hash"`
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func NewUserObject(id, userStatus, employeeid, password_hash string) *User {
	return &User {
		ID:          id,
		User_status: userStatus,
		Employee_id: employeeid,
		Password_hash: password_hash,
	}
}

func (s *UserService) CreateUsersTable() error {
	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		uid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_status VARCHAR(30) NOT NULL,
		employee_id VARCHAR(30) NOT NULL,
        password_hash TEXT NOT NULL
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
	VALUES ($1, $2)
	RETURNING uid;
	`

	err := s.db.QueryRow(insertUserQuery, user.User_status, user.Employee_id, user.Password_hash).Scan(&user.ID)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err);
	}
	log.Println("User created successfully with ID:", user.ID)
	return nil
}
