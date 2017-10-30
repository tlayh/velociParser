package main

import (
	"./service"
	"fmt"
	"os"
	"strings"
)

func main() {

	// if there is a commandline argument, check track name for argument
	// otherwise parse all configured boards

	parseFilter := "false"
	if os.Args[1] != "" {
		parseFilter = os.Args[1]
	}

	config := service.ReadConfig()
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

