package createItem_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/createItem"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateItem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateItem Suite")
}

func TestInputToItemReturnsItemWitCorrectValues(t *testing.T) {
	input := createItem.NewInput("bag")

	item, _ := input.ToItem()

	assert.Equal(t, input.Name, item.Name)
	assert.NotNil(t, item.Id)
}
