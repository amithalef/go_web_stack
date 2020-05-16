package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Item struct {
	Id   uuid.UUID
	Name string
}

func NewItem(name string) (*Item, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return &Item{}, errors.New(fmt.Sprintf("failed to create uuid : %s", err.Error()))
	}
	return &Item{Id: id, Name: name}, nil
}
