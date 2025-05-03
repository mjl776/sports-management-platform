package employees

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type TeamEmployees struct {
	EmployeeID   string `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	EmployeeTitle string `json:"employee_title"`
	SalaryPerHour string `json:"salary_per_hour"`
	EmployerID string `json:"employer_id"`
}

type TeamEmployeesService struct {
	db *sql.DB
}

func NewTeamEmployeesService(db *sql.DB) *TeamEmployeesService {
	return &TeamEmployeesService{
		db: db,
	}
}

func NewTeamEmployeesObject(employeeName, employeeTitle, salaryPerHour, employerID string) *TeamEmployees {
	return &TeamEmployees{
		EmployeeName: "t" + employeeName, // preface employeeID with a 't' to indicate team employee
		EmployeeTitle: employeeTitle,
		SalaryPerHour: salaryPerHour,
		EmployerID: employerID,
	}
}

func (s *TeamEmployeesService) CreateTeamsEmployeesTable() error {
	createEmployeesTableQuery := `
	CREATE TABLE IF NOT EXISTS team_employees (
		employee_id VARCHAR(30) PRIMARY KEY,
		employee_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		employee_title VARCHAR(50) NOT NULL,
		salary_per_hour DECIMAL(10, 2) NOT NULL,
		employer_id INT NOT NULL,
		FOREIGN KEY (employer_id) REFERENCES teams(id),
		CHECK (employee_id ~ '^t[0-9]+$')
	)`

	_, err := s.db.Exec(createEmployeesTableQuery)
	if err != nil {
		return err
	}
	log.Println("Team Employees table created successfully!")
	return nil
}

func (s *TeamEmployeesService) CreateEmployee(employee TeamEmployees) error {
	insertEmployeeQuery := `
	INSERT INTO employees (employee_id, employee_name, created_at, updated_at, employee_title, salary_per_hour, employer_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING employee_id;`

	currentTimestamp := time.Now().Format("2006-01-02 15:04:05")
	employeeID := generateRandomEmployeeID()

	err := s.db.QueryRow(insertEmployeeQuery, employeeID, employee.EmployeeName,
		currentTimestamp, currentTimestamp, employee.EmployeeTitle, employee.SalaryPerHour, employee.EmployerID).Scan(&employeeID)
	if err != nil {
		return err
	}
	return nil
}

func generateRandomEmployeeID() string {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return fmt.Sprintf("%06d", r.Intn(1000000)) // Generates a 6-digit random number
}