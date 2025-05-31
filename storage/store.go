package storage

import (
	"fmt"
	"log"
	"math/rand"
	"quoteService/types"
	"sync"
	"time"
)

var (
	quotes   = []types.Quote{}
	nextID   = 1
	quotesMu = sync.Mutex{}
)

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) GetQuotes() ([]types.Quote, error) {
	quotesMu.Lock()
	defer quotesMu.Unlock()
	return quotes, nil
}

func (a *Adapter) GetRandomQuote() (types.Quote, error) {
	quotesMu.Lock()
	defer quotesMu.Unlock()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rng.Intn(len(quotes))

	return quotes[randomIndex], nil
}

func (a *Adapter) GetQuoteByAuthor(author string) ([]types.Quote, error) {
	quotesMu.Lock()
	defer quotesMu.Unlock()

	var authorQuotes []types.Quote

	for _, q := range quotes {
		if q.Author == author {
			authorQuotes = append(authorQuotes, q)
		}
	}

	if len(authorQuotes) == 0 {
		return nil, fmt.Errorf("no quotes found for author: %s", author)
	}

	return authorQuotes, nil
}

func (a *Adapter) PostQuote(quote types.Quote) (int, error) {
	quotesMu.Lock()
	defer quotesMu.Unlock()
	quote.Id = nextID
	nextID++
	quotes = append(quotes, quote)
	log.Printf("Quotes after adding: %+v", quotes)

	return quote.Id, nil
}

func (a *Adapter) DeleteQuote(id int) error {
	quotesMu.Lock()
	defer quotesMu.Unlock()

	log.Printf("Trying to delete quote with ID: %d", id)

	for i, q := range quotes {
		if q.Id == id {
			quotes = append(quotes[:i], quotes[i+1:]...)
			log.Printf("Deleted quote with ID %d", id)
			return nil
		}
	}
	return fmt.Errorf("quote with ID %d not found", id)
}
