package http

import (
	"encoding/json"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/createItem"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/mongo_storage"
	"net/http"
)

type ItemHttpHandler struct {
	itemStorage mongo_storage.MongoItemStorage
}

func (ih *ItemHttpHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var createItemInput createItem.Input
	if err := json.NewDecoder(r.Body).Decode(&createItemInput); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	//createItem.Usecase{ih.itemStorage}
}
