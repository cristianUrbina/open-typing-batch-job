package domain

type CodeSnippet struct {
	Name       string `json:"name"`
	Content    string `json:"content"`
	Language   string `json:"language"`
	Repository string `json:"repository"`
	RepoDir    string `json:"repo_dir"`
}

func (c CodeSnippet) Equal(other CodeSnippet) bool {
	if c.Name != other.Name || c.Language != other.Language || c.Repository != other.Repository || c.RepoDir != other.RepoDir {
		return false
	}
	return true
}
