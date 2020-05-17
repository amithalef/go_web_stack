package inmemory_storage_test

import (
	str "github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/inmemory_storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewInMemoryStorageReturnStorageWithSizeAsZero(t *testing.T) {
	storage := str.NewInMemoryStorage()
	assert.Equal(t, 0, storage.Size())
}

//func TestShouldBeAbleToAddItem(){
//
//}
