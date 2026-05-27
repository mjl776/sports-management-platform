package employees

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mjl776/sports-management-platform/internal/utils"
)

type TeamEmployees struct {
	ID int  `json:"id"`
	EmployeeID   string `json:"employee_id"`
	EmployeeName string `json:"employee_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	EmployeeTitle string `json:"employee_title"`
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

func NewTeamEmployeesObject(employeeName string, employeeTitle string, employerID string) *TeamEmployees {
	employeeID := util.GenerateRandomULID()

	return &TeamEmployees{
		EmployeeID: employeeID.String(),
		EmployeeName: employeeName,
		EmployeeTitle: employeeTitle,
		EmployerID: employerID,
	}
}

func (s *TeamEmployeesService) CreateTeamsEmployeesTable() error {
	createEmployeesTableQuery := `
	CREATE TABLE IF NOT EXISTS team_employees (
	    id SERIAL PRIMARY KEY,
		employee_id VARCHAR(26) NOT NULL UNIQUE,
		employee_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		employee_title VARCHAR(50) NOT NULL,
		employer_id VARCHAR(26),
		FOREIGN KEY (employer_id) REFERENCES teams(team_id)
	);
`
	_, err := s.db.Exec(createEmployeesTableQuery)
	if err != nil {
		return err
	}
	log.Println("Team Employees table created successfully!")
	return nil
}

func (s *TeamEmployeesService) CreateEmployee(employee TeamEmployees) error {
	// Print current employee
	fmt.Println("Creating employee with name:", employee.EmployeeName,
	 "and title:", employee.EmployeeTitle,
	 "and employer ID:", employee.EmployerID)

	insertEmployeeQuery := `
	INSERT INTO team_employees (employee_id, employee_name, created_at, updated_at, employee_title, employer_id)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;`

	currentTimestamp := time.Now().Format("2006-01-02 15:04:05")

	err := s.db.QueryRow(insertEmployeeQuery, employee.EmployeeID, employee.EmployeeName,
		currentTimestamp, currentTimestamp, employee.EmployeeTitle, employee.EmployerID).Scan(&employee.ID)
	if err != nil {
		return err
	}
	return nil
}
