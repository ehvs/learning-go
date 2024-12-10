package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Words struct {
	//attributes
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: http-get-api <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Println("Invalid URL")
		os.Exit(1)
	}

	response, err := http.Get(args[1])

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("URL is valid: %v \n", response)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output: %d\n : %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	// Unmarshall requires to be a pointer ==> '&words'
	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}
	// to run it, run the test-server.go first, and then 'go run main http://localhost:8080/words'
	//fmt.Printf("HTTP Status Code: %d\nBody: %s\n", response.StatusCode, string(body))
	fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))

}
