package api

import (
	"github.com/gorilla/mux"
)

// Routes return the Router
func (s *Server) Routes() *mux.Router {
	router := mux.NewRouter()

	router.Use(s.loggingMiddleware, s.addRequestIDMiddleware)
	router.HandleFunc("/health", s.health)
	router.HandleFunc("/character", s.list).Methods("GET")
	router.HandleFunc("/character/{id}", s.update).Methods("PUT")
	router.HandleFunc("/character", s.delete).Methods("DELETE")
	router.HandleFunc("/character", s.create).Methods("POST")

	return router
}
