package domain

type CodeSnippet struct {
	Content    []byte
	Language   string
	Repository string
	RepoDir    string
}
