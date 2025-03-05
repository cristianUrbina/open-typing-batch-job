package domain

import "regexp"

func NewFileFilter(ext []string) *FileFilter {
	return &FileFilter{
		ext: ext,
	}
}

type FileFilter struct {
	ext []string
}

func (f *FileFilter) Filter(files []string) ([]string, error) {
	extsRegex := ""
	for _, i := range f.ext {
		extsRegex += i + "|"
	}
	extsRegex = extsRegex[:len(extsRegex)-1]
	return filterStrings(`^.*\.`+extsRegex+"$", files)
}

func filterStrings(pattern string, input []string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	var filtered []string
	for _, str := range input {
		if re.MatchString(str) {
			filtered = append(filtered, str)
		}
	}

	return filtered, nil
}
