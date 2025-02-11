package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type GitHubSearchReposResp struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []struct {
		FullName string `json:"full_name"`
	} `json:"items"`
}

func searchGitHubRepos(lang string) GitHubSearchReposResp {
	baseURL, err := url.Parse("https://api.github.com/search/repositories")
	if err != nil {
		log.Fatal(err)
	}

	values := baseURL.Query()
	values.Add("q", "language:"+lang)
	values.Add("sort", "forks")
	values.Add("order", "desc")
	baseURL.RawQuery = values.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	result, err := bodyToClass[GitHubSearchReposResp](body)
	return result
}

func getRepoTarball(repo string) (io.Reader, error) {
	baseUrl := "https://api.github.com/repos/%s/tarball"
	finalUrl := fmt.Sprintf(baseUrl, repo)
	resp, err := http.Get(finalUrl)
	if err != nil {
		log.Fatalf("Error getting repo tarball: %v", err)
		return nil, err
	}
	return resp.Body, nil
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
