package types

type ServiceQuote interface {
	GetQuotes() ([]Quote, error)
	GetRandomQuote() (Quote, error)
	GetQuoteByAuthor(author string) ([]Quote, error)
	PostQuote(quote Quote) (int, error)
	DeleteQuote(id int) error
}

type StorageQuote interface {
	GetQuotes() ([]Quote, error)
	GetRandomQuote() (Quote, error)
	GetQuoteByAuthor(author string) ([]Quote, error)
	PostQuote(quote Quote) (int, error)
	DeleteQuote(id int) error
}

type Quote struct {
	Id     int    `json:"id"`
	Text   string `json:"quote"`
	Author string `json:"author"`
}
