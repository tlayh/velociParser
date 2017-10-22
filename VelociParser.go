package main

import (
	"fmt"
	"./service"
)

func main() {
	config := service.ReadConfig()
	fmt.Println(config)
	bodyContent := service.ReadLeaderBoard(config.Scene.Url)
	service.ParseLeaderBoardResponse(bodyContent, config.User.Name)
}

