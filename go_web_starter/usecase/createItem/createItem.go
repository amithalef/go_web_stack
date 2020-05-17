package createItem

import (
	"errors"
	"github.com/amithnair91/go_web_stack/go_web_starter/domain"
	"github.com/amithnair91/go_web_stack/go_web_starter/usecase/storage"
)

type Usecase struct {
	ItemStorage storage.ItemStorage
}

type Input struct {
	Name string
	Id   string
}

func (i *Input) ToItem() (*domain.Item, error) {
	return domain.NewItem(i.Name)
}

func (u *Usecase) Execute(input Input) error {
	if len(input.Name) < 1 {
		return errors.New("name cannot be empty")
	}

	exists := u.ItemStorage.Exists(input.Id)
	if exists {
		return errors.New("Item Already exists")
	}
	item, error := input.ToItem()
	if error != nil {
		return error
	}
	u.ItemStorage.Save(item)
	return nil
}
