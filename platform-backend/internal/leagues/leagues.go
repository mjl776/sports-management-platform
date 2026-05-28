package leagues

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"

    "github.com/mjl776/sports-management-platform/internal/utils"
)


type League struct {
	ID int  `json:"id"`
	Name string `json:"name"`
	Sport string `json:"sport"`
	LeagueID string `json:"league_id"`
}

type LeagueService struct {
	db *sql.DB
}

func NewLeagueService(db *sql.DB) *LeagueService {
	return &LeagueService{
		db: db,
	}
}

func NewLeagueObject(name, sport string) *League {

	leagueID := util.GenerateRandomULID();

	return &League{
		Name: name,
		Sport: sport,
		LeagueID: leagueID.String(),
	}
}

func (s *LeagueService) CreateLeaguesTable() error {
	createLeaguesTableSQL := `
	CREATE TABLE IF NOT EXISTS leagues (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		sport VARCHAR(100) NOT NULL,
		league_id VARCHAR(26) NOT NULL UNIQUE
	);
	`

	_, err := s.db.Exec(createLeaguesTableSQL)
	if err != nil {
		log.Fatalf("Failed to create leagues table: %v", err)
	}
	log.Println("Leagues table created successfully!")
	return nil
}

func (s *LeagueService) CreateLeague(league League) error {
	insertLeagueSQL := `
	INSERT INTO leagues (league_id, name, sport)
	VALUES ($1, $2, $3)
	RETURNING id;
`

	err := s.db.QueryRow(insertLeagueSQL, league.LeagueID, league.Name, league.Sport).Scan(&league.ID)
	if err != nil {
		log.Fatalf("Failed to create league: %v", err)
	}
	log.Println("League created successfully!")
	return nil
}
