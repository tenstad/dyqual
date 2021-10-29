package dyqual

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestStruct struct {
	A string
	B string
	C string
	D string
	E string
}

func TestDyqual(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dyqual Suite")
}

var _ = Describe("Dyqual", func() {
	It("dyffs correctly", func() {
		a := TestStruct{
			A: "a",
			B: "b",
			C: "c",
			D: "d",
			E: "e",
		}
		b := TestStruct{
			A: "a",
			B: "b",
			C: "q",
			D: "d",
		}
		Expect(a).To(Dyqual(b))
	})
})
