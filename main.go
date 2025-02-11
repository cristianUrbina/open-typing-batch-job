package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)


func main() {
  fmt.Print("Hello world!\n")
  searchResp := searchGitHubRepos("go")
  fmt.Printf("Parse Search Response: %+v\n", searchResp)
}

func jsonFormatter(s []byte) string {
  var prettyJSON bytes.Buffer
  err := json.Indent(&prettyJSON, s, "", " ")
  if err != nil {
    log.Fatalf("Error formatting JSON: %v", err)
  }
  return prettyJSON.String()
}
