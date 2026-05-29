package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mjl776/sports-management-platform/internal/employees"
	"github.com/mjl776/sports-management-platform/internal/leagues"
	"github.com/mjl776/sports-management-platform/internal/teams"
	"github.com/mjl776/sports-management-platform/internal/users"
	"github.com/mjl776/sports-management-platform/internal/players"
	"github.com/mjl776/sports-management-platform/internal/player-contracts"
)

type APIServer struct {
	listenAddr     string
	leagueService *leagues.LeagueService
	teamsService *teams.TeamsService
	teamEmployeesService *employees.TeamEmployeesService
	usersService *users.UserService
	playersService *players.PlayerService
	playerContractService *playerContracts.PlayerContractService
}

type CreateTeamReqObject struct {
	Name string `json:"name"`
	LeagueID string `json:"league_id"`
}

type CreateLeagueReqObject struct {
	Name string `json:"name"`
	Sport string `json:"sport"`
}

type CreateTeamEmployeeReqObject struct {
	EmployeeName string `json:"employee_name"`
	EmployeeTitle string `json:"employee_title"`
	EmployerID string `json:"employer_id"`
}

type CreateUserReqObject struct {
	UserStatus string `json:"user_status"`
	EmployeeId string `json:"employee_id"`
	Password string `json:"password"`
}

type CreatePlayerReqObject struct {
	Name string `json:"name"`
	TeamID string `json:"team_id"`
}

type CreatePlayerContractObject struct {
	PlayerID string `json:"player_id"`
	TeamID string `json:"team_id"`
	ContractType string `json:"contract_type"`
	ContractLength int `json:"contract_length"`
	Salary float64 `json:"salary"`
}

type LoginReqObject struct {
	EmployeeId string `json:"employee_id"`
	Password string `json:"password"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

func NewAPIServer(listenAddr string,
		leaguesService *leagues.LeagueService,
		teamsService *teams.TeamsService,
		teamEmployeesService *employees.TeamEmployeesService,
		usersService *users.UserService,
		playersService *players.PlayerService,
		playerContractService *playerContracts.PlayerContractService,
	) *APIServer {
	return &APIServer{
		listenAddr:     listenAddr,
		leagueService: leaguesService,
		teamEmployeesService: teamEmployeesService,
		teamsService: teamsService,
		usersService: usersService,
		playersService: playersService,
		playerContractService: playerContractService,
	}
}

func (s *APIServer) Run() {

	router := gin.Default()
	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/create-team", s.handleCreateTeam)
	router.POST("/create-league", s.handleCreateLeague)
	router.POST("/create-team-employee", s.handleCreateTeamEmployee)
	router.POST("/create-user", s.handleCreateUser)
	router.POST("/create-player", s.handleCreatePlayer)
	router.POST("/create-player-contract", s.handleCreatePlayerContract)
	router.POST("/login", s.handleLogin)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create team"})
		return
	}

	c.IndentedJSON(http.StatusOK, team)
}

func (s *APIServer) handleCreateLeague(c *gin.Context) {

	var req CreateLeagueReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	league := leagues.NewLeagueObject(req.Name, req.Sport)
	err := s.leagueService.CreateLeague(*league)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create league"})
		return
	}

	c.IndentedJSON(http.StatusOK, league)
}

func (s *APIServer) handleCreateTeamEmployee(c *gin.Context) {
	var req CreateTeamEmployeeReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee := employees.NewTeamEmployeesObject(req.EmployeeName, req.EmployeeTitle, req.EmployerID)
	err := s.teamEmployeesService.CreateEmployee(*employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create employee."})
		return
	}

	c.IndentedJSON(http.StatusOK, employee)
}

func (s *APIServer) handleCreateUser(c *gin.Context) {
	var req CreateUserReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UserStatus: %s, EmployeeID: %s, Password: %s", req.UserStatus, req.EmployeeId, req.Password)

	user := users.NewUserObject(req.UserStatus, req.EmployeeId, req.Password)
	err := s.usersService.CreateUser(*user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user."})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}


func (s *APIServer) handleLogin(c *gin.Context) {
	var req LoginReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Failed to Login": err.Error()})
		return
	}

	token , err := s.usersService.AuthenticationLogin(req.EmployeeId, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{"token": token})

}

func (s *APIServer) handleCreatePlayer(c *gin.Context) {
	var req CreatePlayerReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player := players.NewPlayerObject(req.Name, req.TeamID)
	err := s.playersService.CreatePlayer(*player)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create player."})
		return
	}

	c.IndentedJSON(http.StatusOK, player)
}

func (s *APIServer) handleCreatePlayerContract(c *gin.Context) {
	var req CreatePlayerContractObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playerContract := playerContracts.NewPlayerContractObject(req.PlayerID, req.Salary, req.ContractType, req.ContractLength, req.TeamID)
	err := s.playerContractService.CreatePlayerContract(*playerContract)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create player contract."})
		return
	}

	c.IndentedJSON(http.StatusOK, playerContract);
}