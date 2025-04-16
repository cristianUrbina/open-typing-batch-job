package infrastructure

import (
	"strings"
	"testing"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/testutils"

	samplecodes "cristianUrbina/open-typing-batch-job/testutils/sample_codes"
)

func MakePythonLang() *domain.Language {
	return &domain.Language{
		Name:  "Python",
		Alias: "python",
	}
}

func TestTreeSitterParserParsePython(t *testing.T) {
	// arrange
	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang:   MakePythonLang(),
		},
		Content: []byte(samplecodes.PythonSampleCode),
	}

	expected := []domain.Snippet{
		{
			Content: `def add(a, b):
    return a + b`,
		},
		{
			Content: `def greet(name):
    return f"Hello, {name}!"`,
		},
		{
			Content: `def is_even(num):
    return num % 2 == 0`,
		},
	}
	parser := NewTreeSitterSnippetExtractor()
	// act
	snippets, err := parser.ExtractSnippets(code)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if !testutils.AreSlicesEqual(expected, snippets) {
		t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
	}
}

func MakeJavaLang() *domain.Language {
	return &domain.Language{
		Name:  "Java",
		Alias: "java",
	}
}

func TestTreeSitterParserParseJava(t *testing.T) {
	// arrange
	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang:   MakeJavaLang(),
		},
		Content: []byte(samplecodes.JavaSampleCode),
	}

	expected := []domain.Snippet{
		{
			Content: `public static int subtract(int a, int b) {
        return a - b;
    }`,
		},
		{
			Content: `public static String reverse(String str) {
        return new StringBuilder(str).reverse().toString();
    }`,
		},
		{
			Content: `public static boolean isPrime(int num) {
        if (num <= 1) return false;
        for (int i = 2; i <= Math.sqrt(num); i++) {
            if (num % i == 0) return false;
        }
        return true;
    }`,
		},
	}
	parser := NewTreeSitterSnippetExtractor()
	// act
	snippets, err := parser.ExtractSnippets(code)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if !testutils.AreSlicesEqual(expected, snippets) {
		t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
	}
}

func TestTreeSitterParserParseJavascript(t *testing.T) {
	// arrange
	codeStr := `

	function add(a, b) {
    return a + b;
}

function greet(name) {
    return "Hello, " + name + "!";
}

const isEven = (num) => {
    return num % 2 === 0;
};`

	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang: &domain.Language{
				Name:  "JavaScript",
				Alias: "js",
			},
		},
		Content: []byte(codeStr),
	}

	expected := []domain.Snippet{
		{
			Content: `function add(a, b) {
    return a + b;
}`,
		},
		{
			Content: `function greet(name) {
    return "Hello, " + name + "!";
}`,
		},
	}

	parser := NewTreeSitterSnippetExtractor()
	// act
	snippets, err := parser.ExtractSnippets(code)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if !testutils.AreSlicesEqual(expected, snippets) {
		t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
	}
}

func TestTreeSitterParserParseC(t *testing.T) {
	// arrange
	codeStr := `
	int add(int a, int b) {
    return a + b;
}

void greet(const char *name) {
    printf("Hello, %s!\n", name);
}

bool isEven(int num) {
    return num % 2 == 0;
}`

	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang: &domain.Language{
				Name:  "C",
				Alias: "c",
			},
		},
		Content: []byte(codeStr),
	}

	expected := []domain.Snippet{
		{
			Content: `int add(int a, int b) {
    return a + b;
}`,
		},
		{
			Content: `void greet(const char *name) {
    printf("Hello, %s!\n", name);
}`,
		},
		{
			Content: `bool isEven(int num) {
    return num % 2 == 0;
}`,
		},
	}

	parser := NewTreeSitterSnippetExtractor()
	// act
	snippets, err := parser.ExtractSnippets(code)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if !testutils.AreSlicesEqual(expected, snippets) {
		t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
	}
}

func TestTreeSitterParserParseGo(t *testing.T) {
	// arrange
	codeStr := `

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet() string {
	return "Hello, " + p.Name + "!"
}

func add(a int, b int) int {
	return a + b
}`

	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang: &domain.Language{
				Name:  "Go",
				Alias: "go",
			},
		},
		Content: []byte(codeStr),
	}

	expected := []domain.Snippet{
		{
			Content: `func (p *Person) Greet() string {
	return "Hello, " + p.Name + "!"
}`,
		},
		{
			Content: `func add(a int, b int) int {
	return a + b
}`,
		},
	}

	parser := NewTreeSitterSnippetExtractor()
	// act
	snippets, err := parser.ExtractSnippets(code)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}
	if !testutils.AreSlicesEqual(expected, snippets) {
		t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
	}
}

func TestTreeSitterParserParseRust(t *testing.T) {
	// arrange
	codeStr := `
struct Person {
    name: String,
    age: u32,
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}
`

	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang: &domain.Language{
				Name:  "Rust",
				Alias: "rust",
			},
		},
		Content: []byte(codeStr),
	}

	expected := []domain.Snippet{
		{
			Content: `struct Person {
    name: String,
    age: u32,
}`,
		},
// 		{
// 			Content: `fn greet(&self) -> String {
//         format!("Hello, {}!", self.name)
// }`,
// 		},
		{
			Content: `fn add(a: i32, b: i32) -> i32 {
    a + b
}`,
		},
	}

	parser := NewTreeSitterSnippetExtractor()
	// act
	snippets, err := parser.ExtractSnippets(code)
	// assert
	if err != nil {
		t.Errorf("error was expected to be nil, but got %v", err)
	}

	// Normalize whitespaces and line breaks before comparing
	normalizeSnippets(&expected)
	normalizeSnippets(&snippets)

	if !testutils.AreSlicesEqual(expected, snippets) {
		t.Errorf("snippets expected to be %v, but got %v", expected, snippets)
	}
}

// Helper function to normalize the snippets
func normalizeSnippets(snippets *[]domain.Snippet) {
	for i := range *snippets {
		(*snippets)[i].Content = normalizeSnippetContent((*snippets)[i].Content)
	}
}

// Helper function to normalize the content (trim spaces and normalize line breaks)
func normalizeSnippetContent(content string) string {
	content = strings.TrimSpace(content)
	content = strings.ReplaceAll(content, "\r\n", "\n") // Normalize Windows line breaks to Unix
	return content
}

