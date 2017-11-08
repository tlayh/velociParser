package main

import (
	"./service"
	"./models"
	"fmt"
	"strings"
	"flag"
	"github.com/fatih/color"
)

func main() {

	// if there is a commandline argument, check track name for argument
	// otherwise parse all configured boards

	var parseFilter string
	var additionalUser string
	var validateBoards string
	flag.StringVar(&parseFilter, "filter", "false", "Filter only for some tracks")
	flag.StringVar(&additionalUser, "user", "false", "Add an additonal user to compare times")
	flag.StringVar(&validateBoards, "validate", "false", "Check if all leaderboards are in the config")
	flag.Parse()

	config := service.ReadConfig()

	// validate boards or default, parse leaderboard rankings
	if validateBoards == "true" {
		service.ParseRankingPage(config)
	} else {
		if additionalUser != "false" {
			newUser := new(service.User)
			newUser.Name = additionalUser
			newUser.Compare = false
			config.Users = append(config.Users, *newUser)
		}

		var results models.Results

		// iterate over defined scenes and trackes
		for _, scene := range config.Scenes {
			if parseFilter == "false" || strings.Contains(scene.Track, parseFilter) {
				//fmt.Println("Scanning Board: ", scene.Track)
				bodyContent := service.ReadLeaderBoard(scene.Url)
				result := service.ParseLeaderBoardResponse(bodyContent, config.Users, scene)
				results.Results = append(results.Results, result)
				//fmt.Println("---------------------------")
			}
		}

		// print leaderboard optimized
		fmt.Println("Leaderboard optimized")
		for _, res := range results.Results {
			c := color.New(color.FgCyan)
			c.Println(res.Track)

			timeNextPlace := 0.0

			for _, trackResults := range res.TrackResults {

				if trackResults.Searched == false {
					timeNextPlace = trackResults.Time
				} else {
					c := color.New(color.FgGreen)
					if (trackResults.Rank > 50) {
						c = color.New(color.FgRed)
					}

					c.Print("Rank: ")
					c.Print(trackResults.Rank)
					c.Print(" Name: ")
					c.Print(trackResults.Name)
					c.Print(" Time: ")
					c.Print(trackResults.Time)
					c.Print(" Distance next place: ")
					difference := trackResults.Time - timeNextPlace
					c.Print(float64(int(difference * 100)) / 100)
					fmt.Println()
				}


			}
		}
	}

}
