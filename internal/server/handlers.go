package server

import (
	"encoding/json"
	"github.com/Koderbek/link_storage_service/internal/helper"
	"net/http"
	"net/url"
	"strconv"
)

func (s *Server) handleCreateLink() http.HandlerFunc {
	type request struct {
		Url string `json:"url"`
	}

	type response struct {
		Code string `json:"short_code"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		_, err := url.ParseRequestURI(req.Url)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		id, err := s.repo.Create(req.Url)
		if err != nil {
			respondError(w, http.StatusUnprocessableEntity, err)
			return
		}

		respond(w, http.StatusCreated, response{Code: helper.IdToCode(id)})
	}
}

func (s *Server) handleLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("short_code")
		link, err := s.cache.GetAndIncr(code)
		if err != nil {
			respondError(w, http.StatusNotFound, err)
			return
		}

		respond(w, http.StatusOK, link)
	}
}

func (s *Server) handleLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limitParam := r.URL.Query().Get("limit")
		limit, err := strconv.ParseUint(limitParam, 10, 0)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		offsetParam := r.URL.Query().Get("offset")
		offset, err := strconv.ParseUint(offsetParam, 10, 0)
		if err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		links, err := s.repo.Links(uint(limit), uint(offset))
		if err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		var urls []string
		for _, link := range links {
			urls = append(urls, link.Url)
		}

		respond(w, http.StatusOK, urls)
	}
}

func (s *Server) handleDeleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("short_code")
		id := helper.CodeToId(code)
		if err := s.repo.Delete(id); err != nil {
			respondError(w, http.StatusNotFound, err)
			return
		}

		s.cache.Delete(code)
		respond(w, http.StatusNoContent, nil)
	}
}

func (s *Server) handleStatsLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("short_code")
		id := helper.CodeToId(code)
		stats, err := s.repo.Stats(id)
		if err != nil {
			respondError(w, http.StatusNotFound, err)
			return
		}

		stats.Code = helper.IdToCode(stats.Id)
		cacheVisits, ok := s.cache.GetCurVisits(code)
		if ok {
			stats.Visits = cacheVisits
		}

		respond(w, http.StatusOK, stats)
	}
}
