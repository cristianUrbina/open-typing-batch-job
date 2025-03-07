package domain

type Language struct {
	ID           int
	Name         string
	Alias        string
	LogoURL      string
	Extensions   []string
	Capabilities []string
}
