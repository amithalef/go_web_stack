package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Item struct {
	Id   uuid.UUID  `json:"id" bson:"id"`
	Name string     `json:"name" bson:"name"`
}

func NewItem(name string) (*Item, error) {
	if len(name) < 1 {
		return nil, errors.New("name cannot be empty")
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return &Item{}, errors.New(fmt.Sprintf("failed to create uuid : %s", err.Error()))
	}
	return &Item{Id: id, Name: name}, nil
}
