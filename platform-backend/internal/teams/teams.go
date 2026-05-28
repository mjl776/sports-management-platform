package teams

import (
	"database/sql"
	"fmt"
	"log"
    "github.com/mjl776/sports-management-platform/internal/utils"

	_ "github.com/lib/pq"
)

type Team struct {
	ID int  `json:"id"`
	TeamID string `json:"team_id"`
	Name string `json:"name"`
	LeagueID string `json:"league_id"`
}

type TeamsService struct {
	db *sql.DB
}

func NewTeamsService(db *sql.DB) *TeamsService {
	return &TeamsService {
		db: db,
	}
}

func NewTeamObject(name string, leagueId string) *Team {

	teamID := util.GenerateRandomULID()

	return &Team{
		Name: name,
		LeagueID: leagueId,
		TeamID: teamID.String(),
	}
}

func (s *TeamsService) CreateTeamsTable() error {
	createTeamsTableSQL := `
	CREATE TABLE IF NOT EXISTS teams (
		id SERIAL PRIMARY KEY,
		team_id VARCHAR(26) NOT NULL UNIQUE,
		name VARCHAR(100) NOT NULL,
		league_id VARCHAR(26) NOT NULL,
		FOREIGN KEY (league_id) REFERENCES leagues(league_id)
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
	INSERT INTO teams (team_id, name, league_id)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	fmt.Println("Creating team with name:", team.Name, "and league ID:", team.LeagueID)

	// Execute the SQL statement
	err := s.db.QueryRow(insertTeamSQL, team.TeamID, team.Name, team.LeagueID).Scan(&team.ID)
	if err != nil {
		log.Fatalf("Failed to create team: %v", err)
		return err
	}

	log.Printf("Team created successfully with Name: %s", team.Name)
	return nil
}