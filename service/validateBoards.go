package service

import (
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func ParseRankingPage(conf Conf) {
	foundUrls := make(map[string]bool)
	seedUrls := [...]string{
		"https://www.velocidrone.com/leaderboard_by_version/15/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/17/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/22/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/12/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/14/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/21/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/19/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/25/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/3/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/7/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/8/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/13/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/16/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/18/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/20/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/23/1.11",
		"https://www.velocidrone.com/leaderboard_by_version/24/1.11"}

	// Channels
	chUrls := make(chan string)
	chFinished := make(chan bool)

	// Kick off the crawl process (concurrently)
	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished)
	}

	// Subscribe to both channels
	for c := 0; c < len(seedUrls); {
		select {
		case url := <-chUrls:
			if strings.Contains(url, "leaderboard") &&
				strings.Contains(url, "1.11") &&
				!strings.Contains(url, "leaderboard_by_version"){
				foundUrls[url] = true
			}
		case <-chFinished:
			c++
		}
	}

	// We're done! Print the results...
	compareToConfig(conf, foundUrls)

	close(chUrls)
}

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition
	return
}

// Extract all http** links from a given webpage
func crawl(url string, ch chan string, chFinished chan bool) {
	resp, err := http.Get(url)

	defer func() {
		// Notify that we're done after this function
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function returns

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}

			// Make sure the url begines in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}
		}
	}
}

func compareToConfig(conf Conf, foundUrls map[string]bool) {
	fmt.Println("Anzahl Tracks Config:")
	fmt.Println(len(conf.Scenes))
	fmt.Println("Anzahl Tracks Crawler:")
	fmt.Println(len(foundUrls))

	for url, _ := range foundUrls {
		found := false
		for _, confUrl := range conf.Scenes {
			if url == confUrl.Url {
				found = true
				continue
			}
		}

		if found == false {
			fmt.Println("Missing track:")
			fmt.Println(url)
		}
	}
}
