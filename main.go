package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mjl776/sports-management-platform/internal/api"
	"github.com/mjl776/sports-management-platform/internal/leagues"
	"github.com/mjl776/sports-management-platform/internal/teams"
)

func main() {

	const (
        host = "localhost"
        port = 5432
        password = "6HzJdB4T0tfFFXExxRY1"
        db_name = "sports-management-platform-db"
    )

    // Connect to the database
    pqsqlconn := fmt.Sprintf("host=%s port=%d user=mlee password=%s dbname=%s sslmode=disable", host, port, password, db_name)
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


    // Create the leagues table
    if err := leaguesService.CreateLeaguesTable(); err != nil {
        log.Fatalf("Failed to create leagues table: %v", err)
    }

    // Create the teams table
    if err := teamsService.CreateTeamsTable(); err != nil {
        log.Fatalf("Failed to create teams table: %v", err)
    }

	server := api.NewAPIServer(":3000", teamsService, leaguesService);
	server.Run();

}