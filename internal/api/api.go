package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)
type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) DataBaseHandler() {
	const (
		host = "localhost"
		port = 5432
		password = ""
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
}

func (s *APIServer) Run() {

	router := http.NewServeMux();
	s.DataBaseHandler();
	log.Println("JSON API server running on port: ", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
