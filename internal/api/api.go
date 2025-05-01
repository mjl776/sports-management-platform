package api

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

    "github.com/mjl776/sports-management-platform/internal/leagues"
	"github.com/mjl776/sports-management-platform/internal/teams"
)

type APIServer struct {
	listenAddr  string
	teamService *teams.TeamsService
    leaguesService *leagues.LeagueService
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenAddr string, teamService *teams.TeamsService, leagueService *leagues.LeagueService) *APIServer {
	return &APIServer{
		listenAddr:  listenAddr,
		teamService: teamService,
        leaguesService: leagueService,
	}
}

func (s *APIServer) Run() {

	router := http.NewServeMux()
	router.HandleFunc("/create-team", s.handleCreateTeam)
    router.HandleFunc("/create-league", s.handleCreateLeague)
	log.Println("JSON API server running on port: ", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *APIServer) handleCreateTeam(w http.ResponseWriter, r *http.Request) {

	scanner := bufio.NewScanner(os.Stdin)

    if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error during scan:", err)
		return
	}

	fmt.Println("Enter your team name!")
	scanner.Scan()
	teamName := scanner.Text()

	fmt.Println("Enter your league Id!")
	scanner.Scan()
	leagueId := scanner.Text()

	teamId, err := generateSecureRandomID(16)

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error generating ID:", err)
    }

	team := teams.NewTeamObject(teamId, teamName, leagueId)
	err = s.teamService.CreateTeam(*team)

	WriteJSON(w, http.StatusOK, team)
}

func (s *APIServer) handleCreateLeague(w http.ResponseWriter, r *http.Request) {

	scanner := bufio.NewScanner(os.Stdin)

    if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error during scan:", err)
		return
	}

	fmt.Println("Enter your league name!")
	scanner.Scan()
	leagueName := scanner.Text()

	fmt.Println("Enter your sport!")
	scanner.Scan()
	leagueSport := scanner.Text()

	leagueId, err := generateSecureRandomID(16)

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error generating ID:", err)
    }

	league := leagues.NewLeagueObject(leagueId, leagueName, leagueSport);
	err = s.leaguesService.CreateLeague(*league)

	WriteJSON(w, http.StatusOK, league)
}

func generateSecureRandomID(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
