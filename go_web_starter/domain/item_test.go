package domain_test

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldBeAbleToCreateItemWithName(t *testing.T) {
	itemName := "bag"
	item, _ := domain.NewItem(itemName)

	assert.Equal(t, itemName, item.Name)
	assert.NotEmpty(t, item.Id)
}
