package dyqual

import (
	gomegamatchers "github.com/onsi/gomega/matchers"
	"github.com/onsi/gomega/types"
	"github.com/tenstad/dyqual/matchers"
)

func Dyqual(expected interface{}) types.GomegaMatcher {
	return matchers.DyqualMatcher{
		Expected: expected,
		EqualMatcher: gomegamatchers.EqualMatcher{
			Expected: expected,
		},
	}
}
