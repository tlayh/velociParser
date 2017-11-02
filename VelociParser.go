package main

import (
	"./service"
	"fmt"
	"strings"
	"flag"
)

func main() {

	// if there is a commandline argument, check track name for argument
	// otherwise parse all configured boards

	var parseFilter string
	var additionalUser string
	flag.StringVar(&parseFilter, "filter", "false", "Filter only for some tracks")
	flag.StringVar(&additionalUser, "user", "false", "Add an additonal user to compare times")
	flag.Parse()

	config := service.ReadConfig()

	if additionalUser != "false" {
		newUser := new(service.User)
		newUser.Name = additionalUser
		newUser.Compare = false
		config.Users = append(config.Users, *newUser)
	}

	// iterate over defined scenes and trackes
	for _, scene := range config.Scenes {
		if parseFilter == "false" || strings.Contains(scene.Track, parseFilter) {
			fmt.Println("Scanning Board: ", scene.Track)
			bodyContent := service.ReadLeaderBoard(scene.Url)
			service.ParseLeaderBoardResponse(bodyContent, config.Users)
			fmt.Println("---------------------------")
		}
	}

}

