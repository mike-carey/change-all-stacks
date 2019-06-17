package change_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestChange(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Change Suite")
}
