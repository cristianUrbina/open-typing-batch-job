package domain

func NewCodeAnalyzer() *CodeAnalyzer {
  return &CodeAnalyzer{
  }
}

type CodeAnalyzer struct { }

func (c *CodeAnalyzer) Analyze(code string) ([]string, error) {
  return []string {`
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
  `}, nil
}
