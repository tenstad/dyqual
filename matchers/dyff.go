package matchers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/onsi/gomega/types"
	yamlv3 "gopkg.in/yaml.v3"
)

type DyffMatcher struct {
	Expected interface{}
	Matcher  types.GomegaMatcher
}

func NewDyffMatcher(expected interface{}, matcher types.GomegaMatcher) types.GomegaMatcher {
	return DyffMatcher{
		Expected: expected,
		Matcher:  matcher,
	}
}

func (m DyffMatcher) Match(actual interface{}) (success bool, err error) {
	return m.Matcher.Match(actual)
}

func (m DyffMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return m.Matcher.NegatedFailureMessage(actual)
}

func (m DyffMatcher) FailureMessage(actual interface{}) (message string) {
	if reflect.TypeOf(m.Expected) != reflect.TypeOf(actual) {
		return m.Matcher.FailureMessage(actual)
	}

	exp, err := yamlv3.Marshal(m.Expected)
	if err != nil {
		return m.Matcher.FailureMessage(actual)
	}
	act, err := yamlv3.Marshal(actual)
	if err != nil {
		return m.Matcher.FailureMessage(actual)
	}

	diff, err := compare(string(exp), string(act))
	if err != nil {
		return m.Matcher.FailureMessage(actual)
	}

	diff = strings.ReplaceAll(diff, "\n", "\n  ")
	return fmt.Sprintf("%s not as expected\n  %s", reflect.TypeOf(m.Expected), diff)
}
