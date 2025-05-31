package service

import (
	"quoteService/storage"
	"quoteService/types"
	"testing"
)

func TestLayerService_PostQuote(t *testing.T) {
	store := storage.NewAdapter()
	svc := NewLayerService(store)

	id, err := svc.PostQuote(types.Quote{Author: "Tolkien", Text: "Not all who wander are lost"})
	if err != nil {
		t.Fatalf("error posting quote: %v", err)
	}

	if id <= 0 {
		t.Errorf("expected valid ID, got %d", id)
	}
}

func TestLayerService_GetQuoteByAuthor(t *testing.T) {
	store := storage.NewAdapter()
	svc := NewLayerService(store)

	svc.PostQuote(types.Quote{Author: "Einstein", Text: "Imagination is more important than knowledge"})
	svc.PostQuote(types.Quote{Author: "Einstein", Text: "Life is like riding a bicycle"})

	quotes, err := svc.GetQuoteByAuthor("Einstein")
	if err != nil {
		t.Fatalf("error getting quotes: %v", err)
	}

	if len(quotes) != 2 {
		t.Errorf("expected 2 quotes, got %d", len(quotes))
	}
}
