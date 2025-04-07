package teams

import (
	"testing"
)

func TestTeamInitialization(t *testing.T) {
	mockTeam := createNewTeam("testing123", "Denver Broncos", make(map[string]*Player), make(map[string]*Employee), 1000.0)

	if mockTeam.ID != "testing123" {
		t.Errorf("expected ID to be '123', got '%s'", mockTeam.ID)
	}

	if mockTeam.Name != "Denver Broncos" {
		t.Errorf("expected team to be 'Denver Broncos', got '%s'", mockTeam.Name)
	}

	if len(mockTeam.Players) != 0 {
		t.Errorf(
			"expected Players to be 'empty map', got '%v'",
			mockTeam.Players,
		)
	}

	if len(mockTeam.Employees) != 0 {
		t.Errorf(
			"expected employee to be 'empty map', got '%v'",
			mockTeam.Employees,
		)
	}

	if mockTeam.Budget != 1000.0 {
		t.Errorf("expected team to be 'Denver Broncos', got '%f'", mockTeam.Budget)
	}

}