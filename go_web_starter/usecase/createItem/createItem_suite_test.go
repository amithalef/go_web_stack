package createItem_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCreateItem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateItem Suite")
}
