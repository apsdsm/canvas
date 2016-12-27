package tcelltools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTcelltools(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tcelltools Suite")
}
