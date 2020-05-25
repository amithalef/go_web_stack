package storage

import (
	"github.com/google/uuid"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/domain"
)

type ItemStorage interface {
	Save(item *domain.Item) (domain.Item, error)
	Exists(id uuid.UUID) (bool, error)
}
