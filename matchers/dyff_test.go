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
	assertEquals(t, diff(
		testStruct{
			A: "0",
		},
		testStruct{
			A: "1",
		},
	),
		`matchers_test.testStruct not as expected
  
  a
    ± value change
      - 0
      + 1
  
  `)
}

func diff(expected interface{}, actual interface{}) string {
	return Dyqual(expected).FailureMessage(actual)
}

func assertEquals(t *testing.T, actual string, expected string) {
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
