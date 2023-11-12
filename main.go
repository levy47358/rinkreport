package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Rink struct {
	Name string `json:"default"`
}

type GameDay struct {
	Date  string `json:"gameWeek"`
	Games []Game `json:"games"`
}

type Game struct {
	Id        int    `json:"id"`
	StartTime string `json:"startTimeUTC"`
	State     string `json:"gameState"`
	HomeTeam  Team   `json:"homeTeam"`
	AwayTeam  Team   `json:"awayTeam"`
	Rink      Rink   `json:"venue"`
	Period    int    `json:"period"`
	Goals     []Goal `json:"goals"`
	Clock     Clock  `json:"clock"`
}

type Team struct {
	Id          int    `json:"id"`
	Abbreviaton string `json:"abbrev"`
	ShotsOnGoal int    `json:"sog"`
	Goals       int    `json:"score"`
}

type Clock struct {
	TimeRemaining  string `json:"timeRemaining"`
	IsIntermission bool   `json:"inIntermission"`
}

type Goal struct {
	Period int    `json:"period"`
	Time   string `json:"timeInPeriod"`
	Scorer Scorer `json:"name"`
	Team   string `json:"teamAbbrev"`
}

type Scorer struct {
	Name string `json:"default"`
}

func toLocal(dt string) string {
	// Parse a time string
	t, err := time.Parse(time.RFC3339, dt)
	if err != nil {
		log.Fatal(err)
	}

	// Convert to local time
	localTime := t.Local()

	return localTime.Format("3:04 PM")
}

func getPerSuff(period int) string {
	var perSuf string

	switch period {
	case 1:
		perSuf = "st"
	case 2:
		perSuf = "nd"
	case 3:
		perSuf = "rd"
	}

	return perSuf
}

func gameState(game Game) string {
	var gameClock string

	perSuf := getPerSuff(game.Period)

	gameState := game.State

	switch gameState {
	case "LIVE":
		gamePre := "Period"
		if game.Clock.IsIntermission {
			gamePre = "Intermission"
		}
		gameClock = strconv.Itoa(game.Period) + perSuf + " " + gamePre + " " + game.Clock.TimeRemaining
	case "OFF":
		gameClock = "Final"
	case "FINAL":
		gameClock = "Final"
	case "FUT":
		gameClock = toLocal(game.StartTime)
	}
	return gameClock
}

func gameScore(game Game) {
	gameClock := gameState(game)
	gameTable := tablewriter.NewWriter(os.Stdout)
	gameTable.SetHeader([]string{game.Rink.Name, gameClock})
	gameTable.Append([]string{game.AwayTeam.Abbreviaton, strconv.Itoa(game.AwayTeam.Goals)})
	gameTable.Append([]string{game.HomeTeam.Abbreviaton, strconv.Itoa(game.HomeTeam.Goals)})
	gameTable.Render()
}

func gameGoals(game Game) {
	if len(game.Goals) != 0 {

		scoringTable := tablewriter.NewWriter(os.Stdout)
		scoringTable.SetHeader([]string{"Team", "Scorer", "Period", "Time"})

		fmt.Println("Scoring Summary:")
		for _, goal := range game.Goals {
			scoringTable.Append([]string{goal.Team, goal.Scorer.Name, strconv.Itoa(goal.Period), goal.Time})
		}
		scoringTable.Render()
	}
}

func getSchedule(date string) GameDay {
	// get the nhl schedule for a date

	response, err := http.Get("https://api-web.nhle.com/v1/score/" + date)
	if err != nil {
		log.Fatal(err)
	}

	// parse response to return a Schedule object
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject GameDay

	json.Unmarshal(responseData, &responseObject)

	return responseObject
}

func validateDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func main() {
	showTeam := flag.String("team", "", "Show only the score where this team is playing")
	showDate := flag.String("date", "", "Show Scores from a specific date")
	showGoals := flag.Bool("goals", false, "Show scoring summary")
	flag.Parse()

	var schDate string

	if *showDate != "" {
		if !validateDate(*showDate) {
			log.Fatal("Invalid date format. Please use YYYY-MM-DD.")
		} else {
			schDate = *showDate
		}
	} else {
		schDate = "now"
	}

	var gamesToShow []Game

	games := getSchedule(schDate).Games

	if *showTeam != "" {
		for _, game := range games {
			if game.HomeTeam.Abbreviaton == *showTeam || game.AwayTeam.Abbreviaton == *showTeam {
				gamesToShow = append(gamesToShow, game)
			}
		}

	} else {
		gamesToShow = games
	}

	for _, game := range gamesToShow {
		gameScore(game)
		if *showGoals {
			gameGoals(game)
		}

	}

}
