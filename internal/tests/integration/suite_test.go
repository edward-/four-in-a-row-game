package integration_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Four in a row Game Suite")
}
