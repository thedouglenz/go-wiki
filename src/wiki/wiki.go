package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const wikiUrl string = "https://en.wikipedia.org/w/api.php?continue=&action=query&prop=extracts&exintro=&explaintext=&format=json&redirects"

/*
 * Expected structure of Wikipedia API response JSON
 * Response should unmarshal to type WikiData
 */
type WikiData struct {
	Query QueryObj
}

type QueryObj struct {
	Pages map[string]Page
}

type Page struct {
	Title   string
	Extract string
}

/*
 * End structure
 */

func httpGet(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gowiki <search term>")
		fmt.Println("<search term> may include spaces")
		os.Exit(1)
	}
	clargs := os.Args[1:]

	// string searchTerm: all command line args except 1st joined by "_"
	searchTerm := strings.Join(clargs, "_")

	url := fmt.Sprintf("%s&titles=%s", wikiUrl, searchTerm)

	body := httpGet(url)

	var w WikiData
	err := json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}

	fmt.Println()
	if _, ok := w.Query.Pages["-1"]; ok {
		fmt.Println("No articles found.")
	} else {
		for k, _ := range w.Query.Pages {
			title := w.Query.Pages[k].Title
			extract := w.Query.Pages[k].Extract

			fmt.Println(title)
			fmt.Println()
			fmt.Println(extract)
		}
	}
	fmt.Println()
}
