package matchers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/onsi/gomega/matchers"
	yamlv3 "gopkg.in/yaml.v3"
)

type DyqualMatcher struct {
	Expected     interface{}
	EqualMatcher matchers.EqualMatcher
}

func (e DyqualMatcher) Match(actual interface{}) (success bool, err error) {
	return e.EqualMatcher.Match(actual)
}

func (e DyqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return e.EqualMatcher.NegatedFailureMessage(actual)
}

func (e DyqualMatcher) FailureMessage(actual interface{}) (message string) {
	if reflect.TypeOf(e.Expected) != reflect.TypeOf(actual) {
		return e.EqualMatcher.FailureMessage(actual)
	}

	exp, err := yamlv3.Marshal(e.Expected)
	if err != nil {
		return e.EqualMatcher.FailureMessage(actual)
	}
	act, err := yamlv3.Marshal(actual)
	if err != nil {
		return e.EqualMatcher.FailureMessage(actual)
	}

	diff, err := compare(string(exp), string(act))
	if err != nil {
		return e.EqualMatcher.FailureMessage(actual)
	}

	diff = strings.ReplaceAll(diff, "\n", "\n  ")
	return fmt.Sprintf("%s not as expected\n  %s", reflect.TypeOf(e.Expected), diff)
}
