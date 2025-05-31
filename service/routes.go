package service

import (
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"quoteService/types"
	"quoteService/utils"
	"strconv"
)

type Handler struct {
	service types.ServiceQuote
}

func NewHandler(service types.ServiceQuote) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes(router *mux.Router) {
	router.HandleFunc("/quotes", h.GetQuotes).Methods(http.MethodGet)
	router.HandleFunc("/quotes/random", h.GetRandomQuote).Methods(http.MethodGet)
	router.HandleFunc("/quotes", h.PostQuote).Methods(http.MethodPost)
	router.HandleFunc("/quotes/{id}", h.DeleteQuote).Methods(http.MethodDelete)
}

func (h *Handler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author != "" {
		log.Println("Author param received:", author)
		quotes, err := h.service.GetQuoteByAuthor(author)
		if err != nil {
			log.Printf("Failed to get quotes by author %s: %v", author, err)
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, quotes)
		return
	}

	quotes, err := h.service.GetQuotes()
	if err != nil {
		log.Printf("Failed to get quotes: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, quotes)
}

func (h *Handler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.GetRandomQuote()
	if err != nil {
		log.Printf("Failed to get random quote: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, quote)

}

func (h *Handler) PostQuote(w http.ResponseWriter, r *http.Request) {
	var quote types.Quote
	if err := utils.ParseJSON(r, &quote); err != nil {
		log.Printf("Failed to parse Quote JSON: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if quote.Author == "" || quote.Text == "" {
		err := errors.New("author and quote fields must not be empty")
		log.Printf("PostQuote error: %v", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if _, err := h.service.PostQuote(quote); err != nil {
		log.Printf("Failed to create quote: %v", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("New quote added: Author=%s", quote.Author)
	utils.WriteJSON(w, http.StatusCreated, "Success")

}

func (h *Handler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	quoteId, _ := strconv.Atoi(id)

	if err := h.service.DeleteQuote(quoteId); err != nil {
		log.Printf("Failed to delete qoute with ID %d: %v", quoteId, err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Printf("Quote with ID %d deleted successfully", quoteId)
	utils.WriteJSON(w, http.StatusOK, "Success")
}
