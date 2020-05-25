package createItem

import (
	"errors"
	"fmt"
	"github.com/google/uuid"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/storage"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

type Input struct {
	Name string
	id   string
}

func NewInput(name string, id string) *Input {
	return &Input{Name: name, id: id}
}

func (i *Input) Id() (uuid.UUID, error) {
	id, err := uuid.Parse(i.id)
	if err != nil {
		return uuid.Nil, errors.New(fmt.Sprintf("id is not a valid uuid :%s", err.Error()))
	}
	return id, nil
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
	id, error := input.Id()
	if error != nil {
		return errors.New(fmt.Sprintf("invalid input :%s", error.Error()))
	}
	exists, storageError := u.ItemStorage.Exists(id)
	if storageError != nil {
		return errors.New(fmt.Sprintf("storage is down : %s", storageError.Error()))
	}
	if exists {
		return errors.New("Item Already exists")
	}
	u.ItemStorage.Save(item)
	return nil
}
