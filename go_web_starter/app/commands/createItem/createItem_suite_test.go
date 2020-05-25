package createItem_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/createItem"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCreateItem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CreateItem Suite")
}

func TestInputIdreturnsUUIDfromGivenString(t *testing.T) {
	uuid := uuid.New()
	input := createItem.NewInput("bag", uuid.String())

	id, err := input.Id()

	assert.Nil(t, err)
	assert.Equal(t, uuid, id)
}

func TestInputIdErrorIfIdIsNotUUID(t *testing.T) {
	input := createItem.NewInput("bag", "invalid-uuid")

	id, err := input.Id()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "id is not a valid uuid")
	assert.Equal(t, id, uuid.Nil)
}
