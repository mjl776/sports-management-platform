package teams

type Team struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Players map[string]*Player `json:"player"`
	Employees map[string]*Employee `json:"employees"`
	Budget float64 `json:"budget"` // temporary place holder
}

type Player struct {
	ID string `json:"id"`
	Name string `json:"string"`
	// other properties to be determined
}

type Employee struct {
	ID string `json:"id"`
	Name string `json:"string"`
	// other properties to be determined
}

func createNewTeam(ID, Name string, Players map[string]*Player, Employees map[string]*Employee, Budget float64) *Team {
	return &Team{
		ID: ID,
		Name: Name,
		Players: Players,
		Employees: Employees,
		Budget: Budget,
	}
}

func createNewPlayer(ID, Name string) *Player {
	return &Player{
		ID: ID,
		Name: Name,
	}
}

func createNewEmployee(ID, Name string) *Employee {
	return &Employee{
		ID: ID,
		Name: Name,
	}
}