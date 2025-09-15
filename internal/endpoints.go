package nrk

import (
	"fmt"
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

func (c *Client) FetchCommit(owner, repo string) (*Commit, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?per_page=1", owner, repo)

	var response []Commit
	if err := c.DoRequest("GET", url, &response); err != nil {
		return nil, fmt.Errorf("failed to fetch commits: %w", err)
	}
	return &response[0], nil
}

func (c *Client) FetchAllRepos(owner string) ([]Repo, error) {
	url := fmt.Sprintf("https://api.github.com/orgs/%s/repos", owner)

	var response []Repo
	if err := c.DoRequest("GET", url, &response); err != nil {
		return nil, fmt.Errorf("failed to fetch repos: %w", err)
	}

	return response, nil

}

func (c *Client) FetchRepoInfo(owner, repo string) (*Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)

	var response Repo
	if err := c.DoRequest("GET", url, &response); err != nil {
		return nil, fmt.Errorf("failed to fetch repo info: %w", err)
	}
	return &response, nil

}
