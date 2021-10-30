package example_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/tenstad/dyqual"

	v1 "k8s.io/api/core/v1"
)

func TestDyqual(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dyqual Suite")
}

var _ = Describe("Dyqual", func() {
	ubuntu := v1.Container{
		Name: "ubuntu",
	}
	alpine := v1.Container{
		Name:  "alpine",
		Ports: []v1.ContainerPort{{ContainerPort: 80, Name: "http"}},
	}

	It("previews Equal diff", func() {
		Expect(alpine).To(Equal(ubuntu))
	})

	It("previews Dyqual diff", func() {
		Expect(alpine).To(Dyqual(ubuntu))
	})
})
