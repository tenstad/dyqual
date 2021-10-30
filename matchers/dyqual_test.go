package matchers_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/tenstad/dyqual"
)

type testStruct struct {
	A string
	B string
}

func TestDyqual(t *testing.T) {
	expected := testStruct{
		A: "0",
		B: "1",
	}
	actual := testStruct{
		A: "0",
	}
	matcher := Dyqual(expected)
	diff := matcher.FailureMessage(actual)

	assertEquals(t,
		`matchers_test.testStruct not as expected
  
  b
    ± value change
      - 1
      +
  
  `,
		diff)
}

func assertEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Fatal(failureMessage(expected, actual))
	}
}

func failureMessage(expected string, actual string) string {
	return fmt.Sprintf("Actual not equal to expected!\nExpected:\n%s\nActual:\n%s",
		wrapLines(expected),
		wrapLines(actual))
}

func wrapLines(text string) string {
	char := "|"
	return fmt.Sprintf("%s%s%s",
		char,
		strings.ReplaceAll(
			text,
			"\n",
			fmt.Sprintf("%s\n%s", char, char)),
		char)
}
