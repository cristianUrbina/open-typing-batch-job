package domain

type CodeSnippet struct {
	Name       string
	Content    string
	Language   string
	Repository string
	RepoDir    string
}

func (c CodeSnippet) Equal(other CodeSnippet) bool {
	if c.Name != other.Name || c.Language != other.Language || c.Repository != other.Repository || c.RepoDir != other.RepoDir {
		return false
	}
	return true
}
