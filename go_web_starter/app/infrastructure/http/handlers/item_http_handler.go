package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/commands/createItem"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/mongo_storage"
)

type ItemHttpHandler struct {
	ItemStorage mongo_storage.MongoItemStorage
}

func (ih *ItemHttpHandler) CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var createItemInput createItem.Input
	if r.Method != "POST" {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf("%s is not supported method", r.Method)))
		return
	}
	if r.Body == nil {
		w.WriteHeader(400)
		w.Write([]byte(errors.New("body cannot be empty").Error()))
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&createItemInput); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	usecase := createItem.Usecase{ItemStorage: ih.ItemStorage}

	err := usecase.Execute(createItemInput)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
}
