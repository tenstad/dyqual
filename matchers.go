package dyqual

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/tenstad/dyqual/matchers"
)

func Dyqual(expected interface{}) types.GomegaMatcher {
	return matchers.NewDyffMatcher(expected, gomega.Equal(expected))
}
