package players

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/mjl776/sports-management-platform/internal/utils"
)

type Player struct {
	ID int  `json:"id"`
	PlayerID string `json:"player_id"`
	Name string `json:"name"`
	TeamID string `json:"team_id"`
}

type PlayerService struct {
	db *sql.DB
}

func NewPlayerService(db *sql.DB) *PlayerService {
	return &PlayerService{
		db: db,
	}
}

func (s *PlayerService) CreatePlayersTable(db *sql.DB) error {
	createPlayersTableSQL := `
	CREATE TABLE IF NOT EXISTS players (
		id SERIAL PRIMARY KEY,
		player_id VARCHAR(26) NOT NULL UNIQUE,
		name VARCHAR(100) NOT NULL,
		team_id VARCHAR(26) NOT NULL,
		user_id VARCHAR(26),
		FOREIGN KEY (user_id) REFERENCES users(user_id)
	);
	`

	_, err := db.Exec(createPlayersTableSQL)
	if err != nil {
		log.Fatalf("Failed to create players table: %v", err)
	}
	log.Println("Players table created successfully!")
	return nil
}

func (s *PlayerService) CreatePlayer (player Player) error {
	insertPlayerSQL := `
	INSERT INTO players (player_id, name, team_id)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	err := s.db.QueryRow(insertPlayerSQL, player.PlayerID, player.Name, player.TeamID).Scan(&player.ID)
	if err != nil {
		log.Fatalf("Failed to create player: %v", err)
	}
	return nil
}

func NewPlayerObject(name string, teamID string) *Player {

	playerID := util.GenerateRandomULID()

	return &Player{
		Name: name,
		TeamID: teamID,
		PlayerID: playerID.String(),
	}
}