package main

import (
	"fmt"
	"cristianUrbina/open-typing-batch-job/fileutils"
)


func main() {
  fmt.Print("Hello world!\n")
  searchResp := searchGitHubRepos("go")
  fmt.Printf("Parse Search Response: %+v\n", searchResp)
  resp, err := getRepoTarball(searchResp.Items[0].FullName)
  if err != nil {
    return
  }
  fileutils.ExtractTarball(resp)
}
