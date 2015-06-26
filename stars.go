package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	owner    = flag.String("owner", "prometheus", "The repository's owner / GitHub org.")
	repo     = flag.String("repo", "prometheus", "The repository name.")
	token    = flag.String("github-token", "", "The GitHub API token to use.")
	maxPages = flag.Int("max-pages", 0, "The maximum number of star API pages to query for.")
)

type star struct {
	StarredAt time.Time `json:"starred_at"`
}

func main() {
	flag.Parse()

	if *token == "" {
		log.Fatalln("No GitHub token specified. Aborting.")
	}

	client := &http.Client{}

	page := 1
	allStars := []star{}
	for {
		// Get the star data, page-by-page.
		url := fmt.Sprintf("https://api.github.com/repos/%s/%s/stargazers?page=%d&?per_page=30", *owner, *repo, page)
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "token "+*token)
		req.Header.Add("Accept", "application/vnd.github.v3.star+json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		resp.Body.Close()

		// Parse the star data.
		stars := []star{}
		err = json.Unmarshal(body, &stars)
		if len(stars) == 0 || (*maxPages > 0 && page > *maxPages) {
			break
		}

		// Show some progress and save all the stars...
		log.Printf("Page %d; stars on page: %d. Rate-limit remaining requests: %s", len(stars), page, resp.Header["X-Ratelimit-Remaining"][0])
		allStars = append(allStars, stars...)
		page++
	}

	for i, t := range allStars {
		fmt.Printf("%d %d\n", t.StarredAt.Unix(), i)
	}
}
