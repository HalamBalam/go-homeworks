package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"homeworks/hw13/pkg/crawler"
	"homeworks/hw13/pkg/index"
	"net/http"
	"strconv"
)

type Service struct {
	router *mux.Router
	ind    *index.Index
}

func New(router *mux.Router, ind *index.Index) *Service {
	s := Service{router: router, ind: ind}
	s.endpoints()
	return &s
}

func (s *Service) endpoints() {
	s.router.HandleFunc("/search/{query}", s.search).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v1/docs", s.createDoc).Methods(http.MethodPost)
	s.router.HandleFunc("/api/v1/docs/{id}", s.updateDocFull).Methods(http.MethodPut)
	s.router.HandleFunc("/api/v1/docs/{id}", s.updateDocPart).Methods(http.MethodPatch)
	s.router.HandleFunc("/api/v1/docs/{id}", s.deleteDoc).Methods(http.MethodDelete)
}

func (s *Service) search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result := s.ind.Find(mux.Vars(r)["query"])

	json.NewEncoder(w).Encode(result)
}

func (s *Service) createDoc(w http.ResponseWriter, r *http.Request) {
	var d crawler.Document
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var docs []crawler.Document
	docs = append(docs, d)
	s.ind.AddDocs(docs)

	json.NewEncoder(w).Encode(s.ind.LastID)
}

func (s *Service) updateDocFull(w http.ResponseWriter, r *http.Request) {
	var d crawler.Document
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, doc, ok := s.ind.FindDoc(id); ok {
		doc.URL = d.URL
		doc.Title = d.Title
		doc.Body = d.Body
	} else {
		http.Error(w, "Doc is not found", http.StatusNotFound)
	}
}

func (s *Service) updateDocPart(w http.ResponseWriter, r *http.Request) {
	var d crawler.Document
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, doc, ok := s.ind.FindDoc(id); ok {
		if len(d.URL) > 0 {
			doc.URL = d.URL
		}
		if len(d.Title) > 0 {
			doc.Title = d.Title
		}
		if len(d.Body) > 0 {
			doc.Body = d.Body
		}
	} else {
		http.Error(w, "Doc is not found", http.StatusNotFound)
	}
}

func (s *Service) deleteDoc(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if i, _, ok := s.ind.FindDoc(id); ok {
		s.ind.Docs = append(s.ind.Docs[:i], s.ind.Docs[i+1:]...)
	} else {
		http.Error(w, "Doc is not found", http.StatusNotFound)
	}
}
