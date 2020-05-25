package createItem

import (
	"errors"
	"fmt"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/storage"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

type Input struct {
	Name string
}

func NewInput(name string) *Input {
	return &Input{Name: name}
}

func (i *Input) ToItem() (*domain.Item, error) {
	return domain.NewItem(i.Name)
}

type Usecase struct {
	ItemStorage storage.ItemStorage
}

func (u *Usecase) Execute(input Input) error {
	item, inputError := input.ToItem()
	if inputError != nil {
		return inputError
	}
	_, storageError := u.ItemStorage.Save(item)
	if storageError != nil {
		return errors.New(fmt.Sprintf("save to storage failed : %s", storageError.Error()))
	}
	return nil
}
