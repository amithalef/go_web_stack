package handlers_test

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/http/handlers"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/mongo_storage"
	"github.com/amithnair91/go_web_stack/go_web_starter/app/infrastructure/test_utils"
)

func setupTestDatabase() (test_utils.MongoTestContainer, mongo_storage.MongoItemStorage) {
	container, database := test_utils.StartMongoDbForTest()
	itemStorage := mongo_storage.NewMongoItemStorage(database)
	return container, itemStorage
}

func TestCreateItemHandlerSuccess(t *testing.T) {
	container, itemStorage := setupTestDatabase()
	defer container.Stop()

	var jsonStr = []byte(`{"Name":"bag"}`)

	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	httpHandler := handlers.ItemHttpHandler{ItemStorage: itemStorage}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.CreateItemHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateItemHandlerReturns400ForInvalidInput(t *testing.T) {
	req, err := http.NewRequest("POST", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}
	httpHandler := handlers.ItemHttpHandler{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.CreateItemHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	assert.NotNil(t, rr.Body)
	assert.Equal(t, "body cannot be empty", rr.Body.String())
}

func TestCreateItemHandlerReturnsBadRequestForNonPostRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}
	httpHandler := handlers.ItemHttpHandler{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.CreateItemHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	assert.NotNil(t, rr.Body)
	assert.Equal(t, fmt.Sprintf("%s is not supported method", req.Method), rr.Body.String())
}

func TestCreateItemHandlerReturnsInternalServerErrorIfUnableToCreateItem(t *testing.T) {
	var jsonStr = []byte(`{"Name":"bag"}`)
	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	httpHandler := handlers.ItemHttpHandler{}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(httpHandler.CreateItemHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	assert.NotNil(t, rr.Body)
	assert.Equal(t, "", rr.Body.String())
}
