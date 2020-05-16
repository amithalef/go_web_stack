package storage

import "github.com/amithnair91/go_web_stack/go_web_starter/domain"

type ItemStorage interface {
	Save(item domain.Item) domain.Item
}
