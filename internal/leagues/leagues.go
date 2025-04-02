package leagues

import "github.com/mjl776/sports-management-platform/internal/teams"

type SportType struct {
	ID string `json:"id"`
	Category string `json:"category"`
	Sports []Sports `json:"sports"`
}

type Sports struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Leagues []Leagues `json:"leagues"`
}

type Leagues struct {
	ID int  `json:"id"`
	Name string `json:"name"`
	Teams []teams.Team `json:"teams"`
}


