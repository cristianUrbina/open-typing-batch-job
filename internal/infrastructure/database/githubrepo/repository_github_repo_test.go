package githubrepo
//
// import (
// 	"cristianUrbina/open-typing-batch-job/internal/infrastructure/clients/githubapiclient"
// 	"cristianUrbina/open-typing-batch-job/testutils"
// 	"testing"
// )
//
// func TestRepositoryGitHubRepo(t *testing.T) {
//   // arrange
//   lang := "c"
// 	apiClient := githubapiclient.NewAPIClient()
//   repo := NewRepositoryGithubRepo(*apiClient)
//   expected, err := testutils.CreateRepositorySlice()  // act
//   if err != nil {
//     t.Fatalf("unexpected error %v", err)
//   }
//   // act
//   result, err := repo.SearchByLang(lang)
//   // assert
//   if err != nil {
//     t.Errorf("error was expected to be nil, but got %v", err)
//   }
//   if result != nil {
//     t.Errorf("result was not expected to be %v, but got %v", expected, result)
//   }
// }
