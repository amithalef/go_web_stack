package createItem

import (
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/usecase/storage"
)

type Usecase struct {
	ItemStorage storage.ItemStorage
}

type Input struct {
	Name string
}

func (i *Input) ToItem() domain.Item {
	return domain.Item{Name: i.Name}
}

func (u *Usecase) Execute(input Input) {
	u.ItemStorage.Save(input.ToItem())
}
