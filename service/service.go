package service

import (
	"log"
	"quoteService/types"
)

type LayerService struct {
	store types.StorageQuote
}

func NewLayerService(store types.StorageQuote) *LayerService {
	return &LayerService{store: store}
}

func (l *LayerService) GetQuotes() ([]types.Quote, error) {
	quotes, err := l.store.GetQuotes()
	if err != nil {
		return nil, err
	}
	return quotes, nil
}

func (l *LayerService) GetRandomQuote() (types.Quote, error) {
	return l.store.GetRandomQuote()
}

func (l *LayerService) GetQuoteByAuthor(author string) ([]types.Quote, error) {
	return l.store.GetQuoteByAuthor(author)
}

func (l *LayerService) PostQuote(quote types.Quote) (int, error) {
	log.Printf("Creating new quote with author: %s; text: %s", quote.Author, quote.Text)
	return l.store.PostQuote(quote)
}

func (l *LayerService) DeleteQuote(id int) error {
	return l.store.DeleteQuote(id)
}
