package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mjl776/sports-management-platform/internal/employees"
	"github.com/mjl776/sports-management-platform/internal/leagues"
	"github.com/mjl776/sports-management-platform/internal/teams"
	"github.com/mjl776/sports-management-platform/internal/users"
)

type APIServer struct {
	listenAddr     string
	leagueService *leagues.LeagueService
	teamsService *teams.TeamsService
	teamEmployeesService *employees.TeamEmployeesService
	usersService *users.UserService
}

type CreateTeamReqObject struct {
	Name string `json:"name"`
	LeagueID int `json:"league_id"`
}

type CreateLeagueReqObject struct {
	Name string `json:"name"`
	Sport string `json:"sport"`
}

type CreatTeamEmployeeReqObject struct {
	EmployeeName string `json:"employee_name"`
	EmployeeTitle string `json:"employee_title"`
	SalaryPerHour float64 `json:"salary_per_hour"`
	EmployerID int `json:"employer_id"`
}

type CreateUserReqObject struct {
	UserStatus string `json:"user_status"`
	EmployeeId string `json:"employee_id"`
	Password string `json:"password_hash"`
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
	) *APIServer {
	return &APIServer{
		listenAddr:     listenAddr,
		leagueService: leaguesService,
		teamEmployeesService: teamEmployeesService,
		teamsService: teamsService,
		usersService: usersService,
	}
}

func (s *APIServer) Run() {

	router := gin.Default()
	router.POST("/create-team", s.handleCreateTeam)
	router.POST("/create-league", s.handleCreateLeague)
	router.POST("/create-team-employee", s.handleCreateTeamEmployee)
	router.POST("/create-user", s.handleCreateUser)
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

	league := leagues.NewLeagueObject(req.Name, req.Sport)
	err := s.leagueService.CreateLeague(*league)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create league"})
		return
	}

	c.IndentedJSON(http.StatusOK, league)
}

func (s *APIServer) handleCreateTeamEmployee(c *gin.Context) {
	var req CreatTeamEmployeeReqObject
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee := employees.NewTeamEmployeesObject(req.EmployeeName, req.EmployeeTitle, req.SalaryPerHour, req.EmployerID)
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

	log.Printf("UserStatus: %s, EmployeeID: %s", req.UserStatus, req.EmployeeId)

	user := users.NewUserObject(req.UserStatus, req.EmployeeId, req.Password)
	err := s.usersService.CreateUser(*user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user."})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}


// func generateSecureRandomID(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.URLEncoding.EncodeToString(bytes), nil
// }
