package main

import (
	"./service"
	"./models"
	"fmt"
	"strings"
	"flag"
	"github.com/fatih/color"
	"time"
	"sort"
)

func main() {

	// if there is a commandline argument, check track name for argument
	// otherwise parse all configured boards
	startTime := time.Now()

	var parseFilter string
	var additionalUser string
	var validateBoards string
	var orderBy string
	var cache bool
	flag.StringVar(&parseFilter, "filter", "false", "Filter only for some tracks")
	flag.StringVar(&additionalUser, "user", "false", "Add an additonal user to compare times")
	flag.StringVar(&validateBoards, "validate", "false", "Check if all leaderboards are in the config")
	flag.StringVar(&orderBy, "orderBy", "rank", "Order by track oder by rank. Values: track , rank (default)")
	flag.BoolVar(&cache, "cache", true, "Use the cache, if the cache is valid. Disable cache with value false")
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

		//var results models.Results
		chResults := make(chan models.Result)
		chFinished := make(chan bool)

		var results models.Results

		// iterate over defined scenes and trackes and register channels
		for _, scene := range config.Scenes {
			// trigger crawling
			go crawl(parseFilter, scene, chResults, chFinished, config.Users, cache, config.CacheLifeTime)
		}

		// subscribe to channels
		for c := 0; c < len(config.Scenes); {
			select {
			case result := <-chResults:
				results.Results = append(results.Results, result)
			case <-chFinished:
				c++
			}
		}

		// print leaderboard optimized
		fmt.Println("Leaderboard optimized")
		fmt.Println("---------------------")

		// check for correct ordering depending on input value
		if (orderBy == "rank") {
			sort.Slice(results.Results, func(i, j int) bool {
				return results.Results[i].TrackResults[0].Rank < results.Results[j].TrackResults[0].Rank
			})
		} else {
			sort.Slice(results.Results, func(i, j int) bool {
				return results.Results[i].Track < results.Results[j].Track
			})
		}

		// count tracks with rank > 50 and < 50
		tracksGtFifty := 0
		tracksLtFifty := 0
		var averageRank int64
		var trackCount int64

		for _, res := range results.Results {
			c := color.New(color.FgCyan)
			c.Println(res.Track)

			timeNextPlace := 0.0

			for _, trackResults := range res.TrackResults {

				if trackResults.Searched == false {
					timeNextPlace = trackResults.Time
				} else {
					c := color.New(color.FgGreen)
					if trackResults.Rank > config.Rank {
						tracksGtFifty++
						c = color.New(color.FgRed)
					} else {
						tracksLtFifty++
					}

					// for ranks outside of range
					if trackResults.Rank >= 999 {
						c = color.New(color.FgMagenta)
					}

					// for average rank
					if trackResults.Rank != 999 {
						averageRank += trackResults.Rank
						trackCount++
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

		fmt.Println("Configured rank: ", config.Rank)
		fmt.Println("Tracks with a ranking better than configured rank: ", tracksLtFifty)
		fmt.Println("Tracks with a ranking worse than configured rank: ", tracksGtFifty)
		fmt.Println("Average rank on scanned tracks: ", float64(averageRank*100 / trackCount) / 100)
		fmt.Println("Total tracks with ranking: ", trackCount)
		fmt.Println("Total tracks scanned: ", len(results.Results))

	}

	fmt.Println("Total time for parsing: ", time.Since(startTime).Seconds(), "Seconds")

}

func crawl(parseFilter string, scene service.Scene, chResult chan models.Result, chFinished chan bool, users []service.User, cache bool, cacheLifeTime float64) {

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	if parseFilter == "false" || strings.Contains(scene.Track, parseFilter) {
		bodyContent := service.ReadLeaderBoard(scene.Url, scene.Track, cache, cacheLifeTime)
		result := service.ParseLeaderBoardResponse(bodyContent, users, scene)

		chResult <- result
	}
}
