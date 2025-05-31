package storage

import (
	"quoteService/types"
	"testing"
)

func TestPostAndGetQuotes(t *testing.T) {
	a := NewAdapter()

	_, err := a.PostQuote(types.Quote{Author: "Author1", Text: "Quote1"})
	if err != nil {
		t.Fatalf("failed to post quote: %v", err)
	}

	quotes, err := a.GetQuotes()
	if err != nil {
		t.Fatalf("failed to get quotes: %v", err)
	}

	if len(quotes) != 1 || quotes[0].Author != "Author1" {
		t.Errorf("unexpected quote data: %+v", quotes)
	}
}

func TestDeleteQuote(t *testing.T) {
	a := NewAdapter()
	id, _ := a.PostQuote(types.Quote{Author: "Author1", Text: "To be deleted"})

	err := a.DeleteQuote(id)
	if err != nil {
		t.Errorf("expected successful delete, got error: %v", err)
	}

	err = a.DeleteQuote(999) // should fail
	if err == nil {
		t.Errorf("expected error when deleting non-existing quote")
	}
}

func TestGetQuoteByAuthor(t *testing.T) {
	a := NewAdapter()
	a.PostQuote(types.Quote{Author: "AuthorX", Text: "Quote X"})
	a.PostQuote(types.Quote{Author: "AuthorY", Text: "Quote Y"})

	quotes, err := a.GetQuoteByAuthor("AuthorX")
	if err != nil || len(quotes) != 1 {
		t.Errorf("expected 1 quote for AuthorX, got %v, error: %v", len(quotes), err)
	}
}
