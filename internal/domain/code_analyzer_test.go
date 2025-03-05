package domain

import (
	"cristianUrbina/open-typing-batch-job/testutils"
	"testing"
)

func createMeaningfulCodeSnippet () string {
	result := `
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int min(int a, int b, int c) {
    if (a < b && a < c) return a;
    if (b < c) return b;
    return c;
}

int levenshtein_distance(const char *s1, const char *s2) {
    int len1 = strlen(s1);
    int len2 = strlen(s2);

    int **dp = (int **)malloc((len1 + 1) * sizeof(int *));
    for (int i = 0; i <= len1; i++)
        dp[i] = (int *)malloc((len2 + 1) * sizeof(int));

    for (int i = 0; i <= len1; i++) dp[i][0] = i;
    for (int j = 0; j <= len2; j++) dp[0][j] = j;

    for (int i = 1; i <= len1; i++) {
        for (int j = 1; j <= len2; j++) {
            int cost = (s1[i - 1] == s2[j - 1]) ? 0 : 1;
            dp[i][j] = min(
                dp[i - 1][j] + 1,    // Deletion
                dp[i][j - 1] + 1,    // Insertion
                dp[i - 1][j - 1] + cost // Substitution
            );
        }
    }

    int result = dp[len1][len2];

    for (int i = 0; i <= len1; i++) free(dp[i]);
    free(dp);

    return result;
}

int main() {
    char str1[] = "kitten";
    char str2[] = "sitting";

    printf("Levenshtein distance between '%s' and '%s' is %d\n",
           str1, str2, levenshtein_distance(str1, str2));

    return 0;
}
`
  return result
}

func TestCodeAnalyzer(t *testing.T) {
  // arrange
  analyzer := NewCodeAnalyzer()
  expected := []string {`
int levenshtein_distance(const char *s1, const char *s2) {
    int len1 = strlen(s1);
    int len2 = strlen(s2);

    int **dp = (int **)malloc((len1 + 1) * sizeof(int *));
    for (int i = 0; i <= len1; i++)
        dp[i] = (int *)malloc((len2 + 1) * sizeof(int));

    for (int i = 0; i <= len1; i++) dp[i][0] = i;
    for (int j = 0; j <= len2; j++) dp[0][j] = j;

    for (int i = 1; i <= len1; i++) {
        for (int j = 1; j <= len2; j++) {
            int cost = (s1[i - 1] == s2[j - 1]) ? 0 : 1;
            dp[i][j] = min(
                dp[i - 1][j] + 1,    // Deletion
                dp[i][j - 1] + 1,    // Insertion
                dp[i - 1][j - 1] + cost // Substitution
            );
        }
    }

    int result = dp[len1][len2];

    for (int i = 0; i <= len1; i++) free(dp[i]);
    free(dp);

    return result;
}
  `}
  // act
  content := createMeaningfulCodeSnippet()
  snippets, err := analyzer.Analyze(content)
  // assert
  if err != nil {
    t.Errorf("error was expected to be nil but got %v", err)
  }
  if !testutils.AreStrSlicesEqual(expected, snippets) {
    t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
  }
}
