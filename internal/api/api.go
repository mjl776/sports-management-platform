package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjl776/sports-management-platform/internal/leagues"
	"github.com/mjl776/sports-management-platform/internal/teams"
)

type APIServer struct {
	listenAddr     string
	leagueService *leagues.LeagueService
	teamsService *teams.TeamsService
}

type CreateTeamReqObject struct {
	Name string `json:"name"`
	LeagueID int `json:"league_id"`
}

type CreateLeagueReqObject struct {
	Name string `json:"name"`
	Sport string `json:"sport"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenAddr string, leaguesService *leagues.LeagueService, teamsService *teams.TeamsService) *APIServer {
	return &APIServer{
		listenAddr:     listenAddr,
		leagueService: leaguesService,
		teamsService: teamsService,
	}
}

func (s *APIServer) Run() {

	router := gin.Default()
	router.POST("/create-team", s.handleCreateTeam)
	router.POST("/create-league", s.handleCreateLeague)
	log.Println("JSON API server running on port: ", s.listenAddr)
	err := http.ListenAndServe(s.listenAddr, router)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *APIServer) handleCreateTeam(c *gin.Context) {

	var req CreateTeamReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team := teams.NewTeamObject(req.Name, req.LeagueID);
	err := s.teamsService.CreateTeam(*team)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		return
	}

	c.IndentedJSON(http.StatusOK, team)
}

func (s *APIServer) handleCreateLeague(c *gin.Context) {

	var req CreateLeagueReqObject

	league := leagues.NewLeagueObject(req.Name, req.Sport)
	err := s.leagueService.CreateLeague(*league)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create league"})
		return
	}

	c.IndentedJSON(http.StatusOK, league)
}

// func generateSecureRandomID(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(bytes), nil
// }
