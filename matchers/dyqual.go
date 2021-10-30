package matchers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/onsi/gomega/types"
	yamlv3 "gopkg.in/yaml.v3"
)

type DyqualMatcher struct {
	Expected interface{}
	Matcher  types.GomegaMatcher
}

func (dm DyqualMatcher) Match(actual interface{}) (success bool, err error) {
	return dm.Matcher.Match(actual)
}

func (dm DyqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return dm.Matcher.NegatedFailureMessage(actual)
}

func (dm DyqualMatcher) FailureMessage(actual interface{}) (message string) {
	if reflect.TypeOf(dm.Expected) != reflect.TypeOf(actual) {
		return dm.Matcher.FailureMessage(actual)
	}

	exp, err := yamlv3.Marshal(dm.Expected)
	if err != nil {
		return dm.Matcher.FailureMessage(actual)
	}
	act, err := yamlv3.Marshal(actual)
	if err != nil {
		return dm.Matcher.FailureMessage(actual)
	}

	diff, err := compare(string(exp), string(act))
	if err != nil {
		return dm.Matcher.FailureMessage(actual)
	}

	diff = strings.ReplaceAll(diff, "\n", "\n  ")
	return fmt.Sprintf("%s not as expected\n  %s", reflect.TypeOf(dm.Expected), diff)
}
