package main

import (
	"encoding/json"
	"log"
	"fmt"
	"time"
	"net/url"
	"strings"
	"net/http"
)

const IssuesURL = "https://api.github.com/search/issues"

type Movie struct {
	Title  string
	Year   int `json:"relearsed"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

var moves = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	objectToJson()
	fmt.Println("-------------")
	objectToJsonFormat()
}

// json.Unmarshal(data, &titles) 解码;
func objectToJson() {
	data, err := json.Marshal(moves)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func objectToJsonFormat() {
	data, err := json.MarshalIndent(moves, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

// 输入流解码
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed :%s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
