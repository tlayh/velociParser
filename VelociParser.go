package main

import (
	"fmt"
	"./service"
)

func main() {
	config := service.ReadConfig()
	fmt.Println(config)
	// bodyContent := service.ReadLeaderBoard(config.Scenes)
	// service.ParseLeaderBoardResponse(bodyContent, config.Users)
}

