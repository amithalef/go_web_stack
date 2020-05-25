package testbuilders

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

func Item(name string) (*domain.Item, error) {
	if name == "" {
		name = "awesome item"
	}
	return domain.NewItem(name)
}
