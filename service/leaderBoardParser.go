package service

import (
	"fmt"
	"strings"
	"net/http"
	"log"
	"io/ioutil"
)

func ParseLeaderBoardResponse(bodyContent string, username string) {
	cleanString := strings.Replace(bodyContent, " ", "", -1)
	// fmt.Print(cleanString)
	index := strings.LastIndex(cleanString, username)
	if index != -1 {
		line := findTrLine(index, cleanString)
		fmt.Println(line)
	} else {
		fmt.Println("%v not found", username)
	}
}

func findTrLine(index int, cleanString string) (string) {

	var startIndex = 0
	var endIndex = 0

	// find beginning of tr-line
	for i := index; i > 0; i-- {
		if cleanString[i:i+4] == "<tr>" {
			startIndex = i
			fmt.Println("Beginning of Line found")
			fmt.Println(startIndex)
			break
		}
	}
	// find end of tr line
	for j := startIndex; j < startIndex+700; j++ {
		if cleanString[j:j+5] == "</tr>" {
			endIndex = j+5
			fmt.Println("End of Line found")
			fmt.Println(endIndex)
			break
		}
	}

	if startIndex != 0 && endIndex != 0 {
		return cleanString[startIndex:endIndex]
	}
	return ""
}

func ReadLeaderBoard(url string) (string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		bodyContent, err := ioutil.ReadAll(response.Body)
		bodyString := string(bodyContent)
		return bodyString
		if err != nil {
			log.Fatal(err)
		}
	}
	return ""
}
