package domain

type Snippet struct {
	Name    string
	Content string
}

func (f Snippet) Equal(other Snippet) bool {
	if f.Name != other.Name || f.Content != other.Content {
		return false
	}
	return true
}
