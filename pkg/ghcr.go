// File: pkg/ghcr.go
package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// authResponse is used to parse the token from the ghcr.io auth server.
type authResponse struct {
	Token string `json:"token"`
}

// tagsResponse is used to parse the list of tags from the registry API.
type tagsResponse struct {
	Tags []string `json:"tags"`
}

// getGhcrAuthToken performs the first step of authentication to get a temporary token.
func getGhcrAuthToken(repo, username, pat string) (string, error) {
	authURL := fmt.Sprintf("https://ghcr.io/token?scope=repository:%s:pull", repo)

	req, err := http.NewRequest("GET", authURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create auth token request: %w", err)
	}

	req.SetBasicAuth(username, pat)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get auth token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get auth token, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read auth token response body: %w", err)
	}

	var authResp authResponse
	if err := json.Unmarshal(body, &authResp); err != nil {
		return "", fmt.Errorf("failed to parse auth token response: %w", err)
	}

	return authResp.Token, nil
}

// FIX: The function signature now correctly accepts THREE arguments: repo, username, and pat.
func ListTags(repo, username, pat string) ([]string, error) {
	// Step 1: Get the short-lived registry token.
	registryToken, err := getGhcrAuthToken(repo, username, pat)
	if err != nil {
		return nil, err
	}

	// Step 2: Use the registry token to list the tags.
	tagsURL := fmt.Sprintf("https://ghcr.io/v2/%s/tags/list", repo)

	req, err := http.NewRequest("GET", tagsURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create tags list request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+registryToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list tags: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list tags, status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read tags list response body: %w", err)
	}

	var tagsResp tagsResponse
	if err := json.Unmarshal(body, &tagsResp); err != nil {
		return nil, fmt.Errorf("failed to parse tags list response: %w", err)
	}

	return tagsResp.Tags, nil
}
