package inmemory_storage

type InMemoryItemStorage struct {
}

func (s InMemoryItemStorage) Size() int {
	return 0
}

func NewInMemoryStorage() InMemoryItemStorage {
	return InMemoryItemStorage{}
}
