package main

import (
	"log"
	"quoteService/cmd/api"
)

func main() {
	server := api.NewServer(":8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
