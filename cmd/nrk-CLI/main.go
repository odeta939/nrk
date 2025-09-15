package main

import (
	"fmt"
	"log"
	"os"

	nrk "github.com/odeta939/nrk/internal"
)

func main() {

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}
	owner := "nrkno"
	repoName := "terraform-registry"

	client := nrk.NewClient()

	response, err := client.FetchRepoInfo(owner, repoName)
	if err != nil {
		log.Fatalf("Error fetching repo info: %v", err)
	}

	nrk.PrintResponse(*response, "Repository")

	commitResponse, err := client.FetchCommit(owner, repoName)
	if err != nil {
		log.Fatalf("Error fetching commit info: %v", err)
	}

	nrk.PrintResponse(*commitResponse, "Latest Commit")

	allRepos, err := client.FetchAllRepos(owner)
	if err != nil {
		log.Fatalf("Error fetching all repos: %v", err)
	}

	fmt.Println("List of all repositories:")
	for _, repo := range allRepos {
		fmt.Printf("- %s: %s\n", repo.Name, repo.URL)
	}

}
