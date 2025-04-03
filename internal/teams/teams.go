package teams

type Team struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Players map[string]*Player `json:"player"`
	Employees map[string]*Employee `json:"employees"`
	Budget int64 `json:"budget"` // temporary place holder
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

