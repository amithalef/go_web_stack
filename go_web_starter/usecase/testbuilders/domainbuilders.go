package testbuilders

import (
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
)

func NewItem(name string) (*domain.Item,error) {
	if name == "" {
		name = "awesome item"
	}
	return domain.NewItem(name)
}
