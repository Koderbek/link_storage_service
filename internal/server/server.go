package server

import (
	"encoding/json"
	"github.com/Koderbek/link_storage_service/internal/database"
	"net/http"
)

type Server struct {
	repo   *database.Repository
	router *http.ServeMux
}

func NewServer(repo *database.Repository) *Server {
	s := &Server{repo: repo, router: http.NewServeMux()}
	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("POST /links", s.handleCreateLink())
	s.router.HandleFunc("GET /links/{short_code}", s.handleLink())
	s.router.HandleFunc("GET /links", s.handleLinks())
	s.router.HandleFunc("DELETE /links/{short_code}", s.handleDeleteLink())
	s.router.HandleFunc("GET /links/{short_code}/stats", s.handleStatsLink())
}

func respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func respondError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	respond(w, status, map[string]string{"error": err.Error()})
}
