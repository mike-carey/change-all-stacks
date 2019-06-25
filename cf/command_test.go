package cf_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/cf"
)

var _ = Describe("Command", func() {

	var (
		cfhome CFHome
		cf CFCommand
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	BeforeEach(func() {
		cfhome = NewTempCFHome("")

		stdout = bytes.Buffer{}
		stderr = bytes.Buffer{}

		cf = NewCFCommand(cfhome, "echo", &stdout, &stderr)
	})

	AfterEach(func() {
		GinkgoWriter.Write([]byte("Writing stdout: "))
		GinkgoWriter.Write(stdout.Bytes())
		GinkgoWriter.Write([]byte("\n"))

		GinkgoWriter.Write([]byte("Writing stderr: "))
		GinkgoWriter.Write(stderr.Bytes())
		GinkgoWriter.Write([]byte("\n"))
	})

	It("Should run the command", func() {
		err := cf.Execute("arg1", "arg2")
		Expect(err).To(BeNil())

		out := stdout.String()
		out = strings.TrimSpace(out)

		Expect(out).To(Equal("arg1 arg2"))
	})

	It("Should print the command", func() {
		str, err := cf.String("arg1", "arg2")
		Expect(err).To(BeNil())

		Expect(str).To(ContainSubstring("echo arg1 arg2"))
	})

})
