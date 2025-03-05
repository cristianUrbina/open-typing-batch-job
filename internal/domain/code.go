package domain

type Code struct {
	Repository Repository
	RepoDir    string
	Content    []byte
}
