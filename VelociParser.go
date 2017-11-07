package main

import (
	"./service"
	"./models"
	"fmt"
	"strings"
	"flag"
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
				fmt.Println("Scanning Board: ", scene.Track)
				bodyContent := service.ReadLeaderBoard(scene.Url)
				result := service.ParseLeaderBoardResponse(bodyContent, config.Users, scene)
				results.Results = append(results.Results, result)
				fmt.Println("---------------------------")
			}
		}
	}



	// print leaderboard optimized
	// fmt.Println("Leaderboard optimized")
	// fmt.Println(results)

}
