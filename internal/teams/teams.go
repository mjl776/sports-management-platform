package teams

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Team struct {
	Name string `json:"name"`
	LeagueID int `json:"league_id"`
}

type TeamsService struct {
	db *sql.DB
}

func NewTeamsService(db *sql.DB) *TeamsService {
	return &TeamsService {
		db: db,
	}
}

func NewTeamObject(name string, leagueId int) *Team {
	return &Team{
		Name: name,
		LeagueID: leagueId,
	}
}

func (s *TeamsService) CreateTeamsTable() error {
	createTeamsTableSQL := `
	CREATE TABLE IF NOT EXISTS teams (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		league_id INT NOT NULL,
		FOREIGN KEY (league_id) REFERENCES leagues(id)
	);
	`

	// Execute the SQL statement
	_, err := s.db.Exec(createTeamsTableSQL)
	if err != nil {
		log.Fatalf("Failed to create teams table: %v", err)
	}
    log.Println("Teams table created successfully!")
	return nil
}

func (s *TeamsService) CreateTeam(team Team) error {
	insertTeamSQL := `
	INSERT INTO teams (name, league_id)
	VALUES ($1, $2)
	RETURNING id;
	`
	var teamID int
	fmt.Println("Creating team with name:", team.Name, "and league ID:", team.LeagueID)

	// Execute the SQL statement
	err := s.db.QueryRow(insertTeamSQL, team.Name, team.LeagueID).Scan(&teamID)
	if err != nil {
		log.Fatalf("Failed to create team: %v", err)
		return err
	}

	log.Printf("Team created successfully with Name: %s", team.Name)
	return nil
}