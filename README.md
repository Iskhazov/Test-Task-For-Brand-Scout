# Test-Task-For-Brand-Scout

## Setup
1. Clone repository
```sh
git clone 
cd quoteService
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Start application
 ```sh
go build -o bin/quoteService cmd/main.go
```
4. Run tests
```sh
go test ./...
```
## API Endpoints
GET /quotes - Get all quotes.  
GET /quotes/random - Get one random quote.  
GET /quotes?author="Author_name" - Search quotes with certain author.  
POST /quotes - Post a new quote.  
DELETE /quotes/{id} - Delete a quote.  
