package service

import (
	"fmt"
	"strings"
	"net/http"
	"log"
	"io/ioutil"
	"golang.org/x/net/html"
	"github.com/fatih/color"
	"strconv"
)

func ParseLeaderBoardResponse(bodyContent string, users []User) {
	cleanString := strings.Replace(bodyContent, " ", "", -1)
	for _, user := range users {
		index := strings.LastIndex(cleanString, user.Name)
		if index != -1 {
			line, startIndex := findTrLine(index, cleanString)
			if user.Compare {// go back with the start index to find the player before the search
				lineBefore, _ := findTrLine(startIndex-100, cleanString)
				parseLineDataIntoModel(lineBefore)
			}
			parseLineDataIntoModel(line)
		} else {
			c := color.New(color.FgRed)
			c.Println("Player ", user.Name, " not found!")
		}
	}
}

/*
here we have a full tr line containing all data including td elements
first td = rank
second td = time
third td = name
 */
func parseLineDataIntoModel(line string) {

	rLine := strings.NewReader(line)
	nodes := html.NewTokenizer(rLine)

	elementCounter := 0
	for {
		tt := nodes.Next()
		switch {
			case tt == html.ErrorToken:
				return
			case tt == html.StartTagToken:
				t := nodes.Token()

				// opening td found
				if t.Data == "td" {
					tt = nodes.Next()
					if tt == html.TextToken{
						i := nodes.Token()
						switch {
							case elementCounter == 0:
								rank, _ := strconv.ParseInt(i.Data, 10, 64)
								c := color.New(color.FgGreen)
								if rank > 50 {
									c = color.New(color.FgRed)
								}
								c.Print("Rank: ", i.Data)
								elementCounter++
							case elementCounter == 1:
								fmt.Print(" Time: ", i.Data)
								elementCounter++
							case elementCounter == 2:
								fmt.Print(" Name: ", strings.TrimSpace(i.Data))
								tt = nodes.Next()
								elementCounter = 0
								fmt.Println()
								return
						}
					}
				}
		}
	}

}

func findTrLine(index int, cleanString string) (string, int) {

	var startIndex = 0
	var endIndex = 0

	// find beginning of tr-line
	for i := index; i > 0; i-- {
		if cleanString[i:i+4] == "<tr>" {
			startIndex = i
			break
		}
	}
	// find end of tr line
	for j := startIndex; j < startIndex+700; j++ {
		if cleanString[j:j+5] == "</tr>" {
			endIndex = j+5
			break
		}
	}

	if startIndex != 0 && endIndex != 0 {
		return cleanString[startIndex:endIndex], startIndex
	}
	return "", 0
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
