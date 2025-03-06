package infrastructure

import (
	"testing"

	"cristianUrbina/open-typing-batch-job/internal/domain"
	"cristianUrbina/open-typing-batch-job/testutils"

	samplecodes "cristianUrbina/open-typing-batch-job/testutils/sample_codes"
)

func TestTreeSitterParserParsePython(t *testing.T) {
	// arrange
	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang:   "python",
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

func TestTreeSitterParserParseJava(t *testing.T) {
	// arrange
	code := &domain.Code{
		Repository: &domain.Repository{
			Author: "someauthor",
			Name:   "somename",
			Lang:   "java",
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
			Lang:   "javascript",
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
