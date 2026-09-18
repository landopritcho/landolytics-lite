// =====================================================================================
// Incoming JSON Response Schemas from ESPN API
// =====================================================================================
package landolyticslite

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// =====================================================================================
// Structs
// =====================================================================================

type ScoreboardResponse struct {
	League League
	Games  []Game
}

type League struct {
	Name         string     `json:"name"`
	Abbreviation string     `json:"abbreviation"`
	SeasonInfo   SeasonInfo `json:"season"`
}

type SeasonInfo struct {
	Year      int
	StartDate string
	EndDate   string
	Type      string
}

type Game struct {
	ID               string
	Date             string
	Description      string
	ShortDescription string
	IsNeutralSite    bool
	Venue            Venue
	Teams            []Team
	Situation        Situation
	Status           Status
	Broadcast        string
}

type Venue struct {
	Name     string   `json:"fullName"`
	Location Location `json:"address"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Team struct {
	ID               string
	IsHomeTeam       bool
	IsWinner         bool
	Score            int
	Location         string
	Name             string
	Abbreviation     string
	DisplayName      string
	ShortDisplayName string
	PrimaryColor     string
	AlternateColor   string
}

type Situation struct {
	Down                  int      `json:"down"`
	Distance              int      `json:"distance"`
	YardLine              int      `json:"yardLine"`
	DownDistanceText      string   `json:"downDistanceText"`
	ShortDownDistanceText string   `json:"shortDownDistanceText"`
	PossessionText        string   `json:"possessionText"`
	IsRedZone             bool     `json:"isRedZone"`
	HomeTimeouts          int      `json:"homeTimeouts"`
	AwayTimeouts          int      `json:"awayTimeouts"`
	LastPlay              LastPlay `json:"lastPlay"`
}

type LastPlay struct {
	Type       string
	Text       string
	ScoreValue int
	TeamID     string
	Yardage    int
}

type Status struct {
	DisplayClock string
	Period       int
	IsTBDOrFlex  bool
	Text         string
	IsComplete   bool
	Description  string
	Detail       string
	ShortDetail  string
}

// =====================================================================================
// Custom Unmarshal Methods
// =====================================================================================

func (s *ScoreboardResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Leagues []League `json:"leagues"`
		Games   []Game   `json:"events"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw.Leagues) != 1 {
		return fmt.Errorf("Expected 1 league, got %d", len(raw.Leagues))
	}

	s.League = raw.Leagues[0]
	s.Games = raw.Games

	return nil
}

func (s *SeasonInfo) UnmarshalJSON(data []byte) error {
	var raw struct {
		Year      int    `json:"year"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Type      struct {
			Name string `json:"name"`
		} `json:"type"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s.Year = raw.Year
	s.StartDate = raw.StartDate
	s.EndDate = raw.EndDate
	s.Type = raw.Type.Name

	return nil
}

func (g *Game) UnmarshalJSON(data []byte) error {
	var raw struct {
		ID               string `json:"id"`
		Date             string `json:"date"`
		Description      string `json:"name"`
		ShortDescription string `json:"shortName"`
		Competitions     []struct {
			IsNeutralSite bool      `json:"neutralSite"`
			Venue         Venue     `json:"venue"`
			Teams         []Team    `json:"competitors"`
			Situation     Situation `json:"situation"`
			Status        Status    `json:"status"`
			Broadcast     string    `json:"broadcast"`
		} `json:"competitions"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	g.ID = raw.ID
	g.Date = raw.Date
	g.Description = raw.Description
	g.ShortDescription = raw.ShortDescription

	if len(raw.Competitions) != 1 {
		fmt.Errorf(
			"Expected 1 competition for game %s, got %d",
			raw.ID,
			len(raw.Competitions),
		)
	}

	g.IsNeutralSite = raw.Competitions[0].IsNeutralSite
	g.Venue = raw.Competitions[0].Venue
	g.Teams = raw.Competitions[0].Teams
	g.Situation = raw.Competitions[0].Situation
	g.Status = raw.Competitions[0].Status
	g.Broadcast = raw.Competitions[0].Broadcast

	return nil
}

func (t *Team) UnmarshalJSON(data []byte) error {
	var raw struct {
		ID       string `json:"id"`
		HomeAway string `json:"homeAway"`
		Winner   bool   `json:"winner"`
		Score    string `json:"score"`
		Team     struct {
			Location         string `json:"location"`
			Name             string `json:"name"`
			Abbreviation     string `json:"abbreviation"`
			DisplayName      string `json:"displayName"`
			ShortDisplayName string `json:"shortDisplayName"`
			Color            string `json:"color"`
			AlternateColor   string `json:"alternateColor"`
		} `json:"team"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	t.ID = raw.ID
	t.IsHomeTeam = strings.ToUpper(raw.HomeAway) == "HOME"
	t.IsWinner = raw.Winner
	t.Score, _ = strconv.Atoi(raw.Score)
	t.Location = raw.Team.Location
	t.Name = raw.Team.Name
	t.Abbreviation = raw.Team.Abbreviation
	t.DisplayName = raw.Team.DisplayName
	t.ShortDisplayName = raw.Team.ShortDisplayName
	t.PrimaryColor = raw.Team.Color
	t.AlternateColor = raw.Team.AlternateColor

	return nil
}

func (l *LastPlay) UnmarshalJSON(data []byte) error {
	var raw struct {
		Type struct {
			Text         string `json:"text"`
			Abbreviation string `json:"abbreviation"`
		} `json:"type"`
		Text       string `json:"text"`
		ScoreValue int    `json:"scoreValue"`
		Team       struct {
			ID string `json:"id"`
		} `json:"team"`
		Yardage int `json:"statYardage"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	l.Type = raw.Type.Text
	l.Text = raw.Text
	l.ScoreValue = raw.ScoreValue
	l.TeamID = raw.Team.ID
	l.Yardage = raw.Yardage

	return nil
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var raw struct {
		DisplayClock string `json:"displayClock"`
		Period       int    `json:"period"`
		Type         struct {
			Text        string `json:"name"`
			IsComplete  bool   `json:"completed"`
			Description string `json:"description"`
			Detail      string `json:"detail"`
			ShortDetail string `json:"shortDetail"`
		} `json:"type"`
		IsTBDOrFlex bool `json:"isTBDFlex"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	s.DisplayClock = raw.DisplayClock
	s.Period = raw.Period
	s.IsTBDOrFlex = raw.IsTBDOrFlex
	s.Text = raw.Type.Text
	s.IsComplete = raw.Type.IsComplete
	s.Description = raw.Type.Description
	s.Detail = raw.Type.Detail
	s.ShortDetail = raw.Type.ShortDetail

	return nil
}
