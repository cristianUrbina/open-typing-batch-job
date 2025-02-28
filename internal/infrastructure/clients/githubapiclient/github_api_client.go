package githubapiclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type APIClient struct {
	baseURL string
	client  *http.Client
}

func NewAPIClient() *APIClient {
	return &APIClient{
		baseURL: "https://api.github.com",
		client:  &http.Client{},
	}
}

type GitHubSearchReposResp struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		FullName string `json:"full_name"`
	} `json:"items"`
}

func (a *APIClient) SearchGitHubRepos(lang string) (*GitHubSearchReposResp, error) {
	apiToken := os.Getenv("API_TOKEN")
	baseURL, err := url.Parse(a.baseURL + "/search/repositories")
	if err != nil {
		log.Fatal(err)
	}

	values := baseURL.Query()
	values.Add("q", "language:"+lang)
	values.Add("sort", "forks")
	values.Add("order", "desc")
	baseURL.RawQuery = values.Encode()

	// resp, err := http.Get(baseURL.String())
	// if err != nil {
	// 	log.Fatalf("Error making GET request: %v", err)
	// }

	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	result, err := bodyToClass[*GitHubSearchReposResp](body)
	return result, nil
}

func (a *APIClient) GetRepoTarball(repo string) (io.ReadSeeker, error) {
	baseURL := "https://api.github.com/repos/%s/tarball"
	finalURL := fmt.Sprintf(baseURL, repo)

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return nil, err
	}

	apiToken := os.Getenv("API_TOKEN")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error getting repo tarball: %v", err)
		return nil, err
	}

	log.Printf("content length: %v", resp.ContentLength)
	log.Println("Rate Limit Remaining:", resp.Header.Get("X-RateLimit-Remaining"))
	log.Println("Rate Limit Reset:", resp.Header.Get("X-RateLimit-Reset"))

	tmpFile, err := os.CreateTemp("", "tarball-*.tar")
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		tmpFile.Close()
		return nil, err
	}

	if _, err := tmpFile.Seek(0, io.SeekStart); err != nil {
		tmpFile.Close()
		return nil, err
	}
	return tmpFile, nil
}

func bodyToClass[T any](body []byte) (T, error) {
	var result T
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return result, err
	}
	return result, nil
}
