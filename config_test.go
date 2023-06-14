package ezenvconfig_test

import (
	"fmt"
	"os"

	"github.com/problem-company-toolkit/ezenvconfig"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ezenvconfig", func() {
	var (
		testEntry ezenvconfig.Entry
	)

	BeforeEach(func() {
		testEntry = ezenvconfig.Entry{
			Name:    "TEST_ENTRY",
			Aliases: []string{"TEST_ALIAS1", "TEST_ALIAS2"},
			OnNotFound: func() {
				fmt.Println("Not found")
			},
			Default:  "default",
			Optional: true,
		}
	})

	Context("ExtractFromEnv function", func() {
		When("environment variable is set", func() {
			BeforeEach(func() {
				os.Setenv("TEST_ALIAS1", "test_value")
			})

			AfterEach(func() {
				os.Unsetenv("TEST_ALIAS1")
			})

			It("should return the correct value", func() {
				value, err := ezenvconfig.ExtractFromEnv(testEntry)
				Expect(err).NotTo(HaveOccurred())
				Expect(value).To(Equal("test_value"))
			})
		})

		When("environment variable is not set", func() {
			BeforeEach(func() {
				os.Unsetenv("TEST_ALIAS1")
				os.Unsetenv("TEST_ALIAS2")
			})

			It("should return the default value", func() {
				value, err := ezenvconfig.ExtractFromEnv(testEntry)
				Expect(err).NotTo(HaveOccurred())
				Expect(value).To(Equal("default"))
			})

			It("should return an error if no default value", func() {
				testEntry.Default = ""
				value, err := ezenvconfig.ExtractFromEnv(testEntry)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("could not find a value for entry"))
				Expect(value).To(Equal(""))
			})
		})
	})
})
