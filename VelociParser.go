package main

import (
	"./service"
	"fmt"
)

func main() {
	config := service.ReadConfig()
	// iterate over defined scenes and trackes
	for _, scene := range config.Scenes {
		fmt.Println("Scanning Board: ", scene.Track)
		bodyContent := service.ReadLeaderBoard(scene.Url)
		service.ParseLeaderBoardResponse(bodyContent, config.Users)
		fmt.Println("---------------------------")
	}

}

