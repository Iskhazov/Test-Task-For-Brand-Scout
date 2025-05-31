package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"quoteService/service"
	"quoteService/storage"
)

// Server — main server structure, contains address and database connection
type Server struct {
	addr string
}

// NewServer creates and returns a new instance of the server
func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

// Run starts the HTTP server and registers all routes
func (s *Server) Run() error {
	router := mux.NewRouter() // Create a new router using Gorilla Mux
	subrouter := router.PathPrefix("").Subrouter()

	requestStore := storage.NewAdapter()
	requestService := service.NewLayerService(requestStore)
	requestHandler := service.NewHandler(requestService)

	requestHandler.Routes(subrouter) // Register handler routes in the router

	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router) // Start the HTTP server
}
