package createItem

import (
	"errors"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/storage"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

type Input struct {
	Name string
	Id   string
}

func (i *Input) ToItem() (*domain.Item, error) {
	return domain.NewItem(i.Name)
}

type Usecase struct {
	ItemStorage storage.ItemStorage
}

func (u *Usecase) Execute(input Input) error {
	item, error := input.ToItem()
	if error != nil {
		return error
	}
	exists := u.ItemStorage.Exists(input.Id)
	if exists {
		return errors.New("Item Already exists")
	}
	u.ItemStorage.Save(item)
	return nil
}
