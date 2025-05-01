package leagues

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)


type League struct {
	ID string  `json:"id"`
	Name string `json:"name"`
	Sport string `json:"sport"`
}

type LeagueService struct {
	db *sql.DB
}

func NewLeagueService(db *sql.DB) *LeagueService {
	return &LeagueService{
		db: db,
	}
}


func NewLeagueObject(id, name, sport string) *League {
	return &League{
		ID: id,
		Name: name,
		Sport: sport,
	}
}

func (s *LeagueService) CreateLeaguesTable() error {
	createLeaguesTableSQL := `
	CREATE TABLE IF NOT EXISTS leagues (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		sport VARCHAR(100) NOT NULL
	);
	`

	// Execute the SQL statement
	_, err := s.db.Exec(createLeaguesTableSQL)
	if err != nil {
		log.Fatalf("Failed to create leagues table: %v", err)
	}
	log.Println("Leagues table created successfully!")
	return nil
}

func (s *LeagueService) CreateLeague(league League) error {
	insertLeagueSQL := `
	INSERT INTO leagues (name, sport)
	VALUES ($1, $2)
	RETURNING id;
	`

	err := s.db.QueryRow(insertLeagueSQL, league.Name, league.Sport).Scan(&league.ID)
	if err != nil {
		log.Fatalf("Failed to create league: %v", err)
	}
	log.Println("League created successfully!")
	return nil
}

