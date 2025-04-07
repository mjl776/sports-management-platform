package leagues

import "github.com/mjl776/sports-management-platform/internal/teams"

type SportType struct {
	ID string `json:"id"`
	Category string `json:"category"`
	Sports map[string]*Sport `json:"sports"`
}

type Sport struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Leagues map[string]*League `json:"leagues"`
}

type League struct {
	ID int  `json:"id"`
	Name string `json:"name"`
	Teams map[string]*teams.Team `json:"teams"`
}


