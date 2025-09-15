package nrk

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Client struct {
	HTTPClient *http.Client
	Token      string
}

func NewClient() *Client {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}

	return &Client{
		HTTPClient: &http.Client{},
		Token:      token,
	}
}

func (c *Client) DoRequest(method, url string, objectToDecodeTo any) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("User-Agent", "nrk-client")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(objectToDecodeTo); err != nil {
		return err
	}

	return nil
}
