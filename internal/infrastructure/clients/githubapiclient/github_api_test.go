package githubapiclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestSearchGitHubRepos(t *testing.T) {
	// arrange
	expectedJSON := `{
    "total_count": 35225821,
    "incomplete_results": false,
    "items": [
        {
            "name": "bootstrap",
            "full_name": "twbs/bootstrap"
        }
      ]
    }`
	expectedResult, err := bodyToClass[*GitHubSearchReposResp]([]byte(expectedJSON))
	if err != nil {
		t.Fatalf("Failed to parse expectedResult: %v", err)
	}

	lang := "c"
	token := "mysecrettoken"
	err = os.Setenv("API_TOKEN", token)
	if err != nil {
		t.Fatalf("Failed to set environemnt variable: %v", err)
	}

	defer os.Unsetenv("API_TOKEN")
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		apiToken := token
		expectedHeader := "Bearer " + apiToken

		if authHeader != expectedHeader {
			t.Errorf("Expected header %q, but got %q", expectedHeader, authHeader)
		}

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, expectedJSON)
	}))
	defer mockServer.Close()

	// act
	client := &APIClient{
		baseURL: mockServer.URL,
		client:  &http.Client{},
	}
	result, err := client.SearchGitHubRepos(lang)
	// assert
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Expected result: %v, got %v", expectedResult, result)
	}
}
