package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Repo struct {
	Name        string  `json:"name"`
	URL         string  `json:"html_url"`
	Description string  `json:"description"`
	License     License `json:"license"`
}

type License struct {
	Name string `json:"name"`
}

type Commit struct {
	Date   string `json:"date"`
	Commit struct {
		Message string `json:"message"`
		Author  struct {
			Name string    `json:"name"`
			Date time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

func main() {

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}
	owner := "nrkno"
	repoName := "terraform-registry"

	response, err := fetchRepoInfo(owner, repoName, token)
	if err != nil {
		log.Fatalf("Error fetching repo info: %v", err)
	}

	fmt.Println("===============================================")
	fmt.Println("Repository Information:")
	fmt.Println("===============================================")

	fmt.Printf("%-12s %s\n", "Name:", response.Name)
	fmt.Printf("%-12s %s\n", "URL:", response.URL)
	fmt.Printf("%-12s %s\n", "Description:", response.Description)
	fmt.Printf("%-12s %s\n", "License:", response.License.Name)
	fmt.Println("===============================================")

	commitResponse, err := fetchCommit(owner, repoName, token)
	if err != nil {
		log.Fatalf("Error fetching commit info: %v", err)
	}

	msg := commitResponse.Commit.Message
	title := strings.SplitN(msg, "\n", 2)[0]

	println()
	fmt.Println("===============================================")
	fmt.Println("Latest Commit Information:")
	fmt.Println("===============================================")
	fmt.Printf("%-12s %s\n", "Author:", commitResponse.Commit.Author.Name)
	fmt.Printf("%-12s %s\n", "Date:", commitResponse.Commit.Author.Date)
	fmt.Printf("%-12s %s\n", "Message:", title)
	fmt.Println("===============================================")

}

func fetchCommit(owner, repo, token string) (*Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: %s, body: %s", resp.Status, string(body))
	}

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read response body: %w", err)
	// }
	// fmt.Println(string(body))

	// var response []Commit
	// if err := json.Unmarshal(body, &response); err != nil {
	// 	return nil, fmt.Errorf("failed to decode JSON: %w", err)
	// }

	var response []Commit
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &response[0], nil

}

func fetchRepoInfo(owner, repo, token string) (*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: %s, body: %s", resp.Status, string(body))
	}

	var response Repo
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return &response, nil

}
