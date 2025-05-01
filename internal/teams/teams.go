package teams

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type TeamsService struct {
	db *sql.DB
}

type Team struct {
	ID string `json:"id"`
	Name string `json:"name"`
	LeagueID string `json:"league_id"`
}

func NewTeamObject(id string, name string, leagueId string) *Team {
	return &Team{
		ID: id,
		Name: name,
		LeagueID: leagueId,
	}
}

func NewTeamsService(db *sql.DB) *TeamsService {
	return &TeamsService{
		db: db,
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

	err := s.db.QueryRow(insertTeamSQL, team.Name, team.LeagueID).Scan(&team.ID)
	if err != nil {
		log.Fatalf("Failed to create team: %v", err)
		return err
	}

	log.Printf("Team created successfully with ID: %s", team.ID)
	return nil
}