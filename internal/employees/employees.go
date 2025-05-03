package employees

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Employees struct {
	EmployeeID   string `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	EmployeeTitle string `json:"employee_title"`
	SalaryPerHour string `json:"salary_per_hour"`
}

type EmployeesService struct {
	db *sql.DB
}

func NewEmployeesService(db *sql.DB) *EmployeesService {
	return &EmployeesService{
		db: db,
	}
}

func NewEmployeesObject(employeeID, employeeName, createdAt, updatedAt, employeeTitle, salaryPerHour string) *Employees {
	return &Employees{
		EmployeeID:   employeeID,
		EmployeeName: employeeName,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		EmployeeTitle: employeeTitle,
		SalaryPerHour: salaryPerHour,
	}
}

func (s *EmployeesService) CreateEmployeesTable() error {
	createEmployeesTableQuery := `
	CREATE TABLE IF NOT EXISTS employees (
		employee_id VARCHAR(30) PRIMARY KEY,
		employee_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		employee_title VARCHAR(50) NOT NULL,
		salary_per_hour DECIMAL(10, 2) NOT NULL,
		CHECK (employee_id ~ '^[0-9]+$')
	)`

	_, err := s.db.Exec(createEmployeesTableQuery)
	if err != nil {
		return err
	}
	log.Println("Leagues table created successfully!")
	return nil
}

func (s *EmployeesService) CreateEmployee(employee Employees) error {
	insertEmployeeQuery := `
	INSERT INTO employees (employee_id, employee_name, created_at, updated_at, employee_title, salary_per_hour)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING employee_id;`

	err := s.db.QueryRow(insertEmployeeQuery, generateRandomEmployeeID(), employee.EmployeeName,
		employee.CreatedAt, employee.UpdatedAt, employee.EmployeeTitle, employee.SalaryPerHour).Scan(&employee.EmployeeID)
	if err != nil {
		return err
	}
	return nil
}

func generateRandomEmployeeID() string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return fmt.Sprintf("%06d", r.Intn(1000000)) // Generates a 6-digit random number
}