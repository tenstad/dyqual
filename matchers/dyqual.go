package dyqual

import (
	"github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	yamlv3 "gopkg.in/yaml.v3"
)

type DyqualMatcher struct {
	Expected     interface{}
	EqualMatcher matchers.EqualMatcher
}

func Dyqual(expected interface{}) types.GomegaMatcher {
	return DyqualMatcher{
		Expected: expected,
		EqualMatcher: matchers.EqualMatcher{
			Expected: expected,
		},
	}
}

func (e DyqualMatcher) Match(actual interface{}) (success bool, err error) {
	return e.EqualMatcher.Match(actual)
}

func (e DyqualMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return e.EqualMatcher.NegatedFailureMessage(actual)
}

func (e DyqualMatcher) FailureMessage(actual interface{}) (message string) {
	exp, expErr := yamlv3.Marshal(e.Expected)
	act, actErr := yamlv3.Marshal(actual)
	if expErr != nil || actErr != nil {
		return e.EqualMatcher.FailureMessage(actual)
	}

	diff, err := compare(string(act), string(exp))
	if err != nil {
		return e.EqualMatcher.FailureMessage(actual)
	}
	return *diff
}
