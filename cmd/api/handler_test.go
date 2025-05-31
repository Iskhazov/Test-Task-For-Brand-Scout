package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"quoteService/service"
	"quoteService/storage"
	"quoteService/types"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	store := storage.NewAdapter()
	svc := service.NewLayerService(store)
	handler := service.NewHandler(svc)

	r := mux.NewRouter()
	handler.Routes(r)
	return r
}

func TestPostQuoteHandler(t *testing.T) {
	router := setupRouter()

	body := types.Quote{
		Author: "Pushkin",
		Text:   "Ya pomnyu chudnoe mgnovenie...",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
}

func TestGetQuoteByAuthorHandler(t *testing.T) {
	router := setupRouter()

	// Добавим цитату
	postReqBody := types.Quote{Author: "Lermontov", Text: "Vykhozhu odin ya na dorogu..."}
	postBody, _ := json.Marshal(postReqBody)
	reqPost, _ := http.NewRequest("POST", "/quotes", bytes.NewBuffer(postBody))
	reqPost.Header.Set("Content-Type", "application/json")
	rrPost := httptest.NewRecorder()
	router.ServeHTTP(rrPost, reqPost)

	// Запрос по автору
	req, _ := http.NewRequest("GET", "/quotes?author=Lermontov", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rr.Code)
	}
}
