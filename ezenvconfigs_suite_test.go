package ezenvconfig_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEzenvconfigs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ezenvconfigs Suite")
}
