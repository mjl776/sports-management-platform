package playerContracts

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PlayerContract struct {
	ID        int     `json:"id"`
	PlayerID  string  `json:"player_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Salary    float64 `json:"salary"`
	ContractType string  `json:"contract_type"`
	ContractLength int `json:"contract_length"`
	TeamID string `json:"team_id"`
}

type PlayerContractService struct {
	db *sql.DB
}

func NewPlayerContractService(db *sql.DB) *PlayerContractService {
	return &PlayerContractService{
		db: db,
	}
}

func (s *PlayerContractService) CreatePlayerContractsTable(db *sql.DB) error {
	createPlayerContractsTableSQL := `
	CREATE TABLE IF NOT EXISTS player_contracts (
		id SERIAL PRIMARY KEY,
		player_id VARCHAR(26) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		salary NUMERIC(15, 2) NOT NULL,
		contract_type VARCHAR(50) NOT NULL,
		contract_length INT NOT NULL,
		team_id VARCHAR(26) NOT NULL
		FOREIGN KEY (player_id) REFERENCES players(player_id),
    	FOREIGN KEY (team_id) REFERENCES teams(team_id)
	);
	`

	_, err := s.db.Exec(createPlayerContractsTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func NewPlayerContractObject(playerID string, salary float64, contractType string, contractLength int, teamID string) *PlayerContract {
	return &PlayerContract{
		PlayerID: playerID,
		Salary: salary,
		ContractType: contractType,
		ContractLength: contractLength,
		TeamID: teamID,
	}
}

func (s *PlayerContractService) CreatePlayerContract(playerContract PlayerContract) error {
	insertPlayerContractSQL := `
	INSERT INTO player_contracts (player_id, salary, contract_type, contract_length, team_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	err := s.db.QueryRow(insertPlayerContractSQL, playerContract.PlayerID, playerContract.Salary, playerContract.ContractType, playerContract.ContractLength, playerContract.TeamID).Scan(&playerContract.ID)
	if err != nil {
		return err
	}
	return nil
}