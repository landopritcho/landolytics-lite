// =====================================================================================
// ESPN Scoreboard Interface
// =====================================================================================
package landolyticslite

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// =====================================================================================
// Constants
// =====================================================================================

const scoreboardURL = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"

// =====================================================================================
// Public Functions
// =====================================================================================

func GetNFLScoreboard(date string) (ScoreboardResponse, error) {
	var scoreboardResponse ScoreboardResponse
	formattedURL := fmt.Sprintf("%s?dates=%s", scoreboardURL, date)

	response, err := http.Get(formattedURL)
	if err != nil {
		return scoreboardResponse, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return scoreboardResponse, err
	}

	err = json.Unmarshal(body, &scoreboardResponse)
	if err != nil {
		return scoreboardResponse, err
	}

	return scoreboardResponse, nil
}
