package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

func TestShouldBeAbleToCreateItemWithName(t *testing.T) {
	itemName := "bag"
	item, _ := domain.NewItem(itemName)

	assert.Equal(t, itemName, item.Name)
	assert.NotEmpty(t, item.Id)
}

func TestNewItemFailsWhenNameIsEmpty(t *testing.T) {
	item, err := domain.NewItem("")

	assert.NotNil(t, err)
	assert.Equal(t, "Name cannot be empty", err.Error())
	assert.Nil(t, item)
}
