package config_test

import (
	"os"
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/change-all-stacks/config"
)

const (
	FoundationName = "foundation"
	ApiUrl = "https://api.sys.mydomain.com"
	Username = "username"
	Password = "password"
	SkipSslValidation = true
)

func checkConfigsReturn(c Configs, e error) {
	Expect(e).To(BeNil())
	Expect(c).Should(HaveKey(FoundationName))
	Expect(c[FoundationName].ApiAddress).To(Equal(ApiUrl))
	Expect(c[FoundationName].Username).To(Equal(Username))
	Expect(c[FoundationName].Password).To(Equal(Password))
	Expect(c[FoundationName].SkipSslValidation).To(Equal(SkipSslValidation))
}

var _ = Describe("Config", func() {

	var (
		json_str string
	)

	BeforeEach(func() {
		json_str = fmt.Sprintf(`{"%s": {"api_url": "%s", "user": "%s", "password": "%s", "skip_ssl_validation": %t}}`, FoundationName, ApiUrl, Username, Password, SkipSslValidation)
	})

	It("Should load a string", func() {
		checkConfigsReturn(LoadConfigFromString(json_str))
	})

	It("Should load bytes", func() {
		checkConfigsReturn(LoadConfigFromBytes([]byte(json_str)))
	})

	Describe("File", func() {
		var (
			tempFile string
		)

		BeforeEach(func() {
			f, e := ioutil.TempFile(os.TempDir(), "")
			Expect(e).To(BeNil())

			tempFile = f.Name()

			e = ioutil.WriteFile(tempFile, []byte(json_str), 0644)
			Expect(e).To(BeNil())
		})

		AfterEach(func() {
			err := os.Remove(tempFile)
			Expect(err).To(BeNil())
		})

		It("Should load a file", func() {
			checkConfigsReturn(LoadConfigFromFile(tempFile))
		})
	})
})
