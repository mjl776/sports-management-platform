package main

import (
	"database/sql"
	"fmt"
	"log"
    "os"

	_ "github.com/lib/pq"
    "github.com/joho/godotenv"
	"github.com/mjl776/sports-management-platform/internal/api"
	"github.com/mjl776/sports-management-platform/internal/employees"
	"github.com/mjl776/sports-management-platform/internal/leagues"
	"github.com/mjl776/sports-management-platform/internal/teams"
	"github.com/mjl776/sports-management-platform/internal/users"
    "github.com/mjl776/sports-management-platform/internal/players"
    "github.com/mjl776/sports-management-platform/internal/player-contracts"
)

func main() {

    // Load environment variables
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    var (
        host = "localhost"
        port = 5432
        password = os.Getenv("DB_PASSWORD")
        db_name = os.Getenv("DB_NAME")
        user = os.Getenv("DB_USER")
    )

    // Connect to the database
    pqsqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)
    db, err := sql.Open("postgres", pqsqlconn)

    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    // Verify the connection
    err = db.Ping()
    if err != nil {
        log.Fatalf("Unable to ping the database: %v", err)
    }

    log.Println("Successfully connected to the database!")

    teamsService := teams.NewTeamsService(db)
    leaguesService := leagues.NewLeagueService(db)
	usersService := users.NewUserService(db)
	teamsEmployeeService := employees.NewTeamEmployeesService(db)
    playersService := players.NewPlayerService(db)
    playerContractService := playerContracts.NewPlayerContractService(db)

    // Create the leagues table
    if err := leaguesService.CreateLeaguesTable(); err != nil {
        log.Fatalf("Failed to create leagues table: %v", err)
    }

    // Create the teams table
    if err := teamsService.CreateTeamsTable(); err != nil {
        log.Fatalf("Failed to create teams table: %v", err)
    }

	// Create the employees table
	if err := teamsEmployeeService.CreateTeamsEmployeesTable(); err != nil {
		log.Fatalf("Failed to create employees table: %v", err)
	}

	// Create the users table
	if err := usersService.CreateUsersTable(); err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

    if err := playersService.CreatePlayersTable(db); err != nil {
        log.Fatalf("Failed to create players table: %v", err)
    }

    if err := playerContractService.CreatePlayerContractsTable(db); err != nil {
        log.Fatalf("Failed to create player contracts table: %v", err)
    }

    server := api.NewAPIServer(":8080",
                leaguesService,
                teamsService,
                teamsEmployeeService,
                usersService,
                playersService,
                playerContractService,
            );
    server.Run();

}