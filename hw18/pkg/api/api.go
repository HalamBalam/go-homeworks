package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strings"
	"sync"
)

type Service struct {
	router *mux.Router
	links  []string
	mu     sync.Mutex
}

func New(r *mux.Router) *Service {
	s := Service{}
	s.router = r
	s.links = make([]string, 0)
	s.endpoints()
	return &s
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = uint64(len(alphabet))
	domain   = "http://localhost/"
)

func (s *Service) endpoints() {
	s.router.HandleFunc("/api/v1/links/make", s.createLink).Methods(http.MethodPost)
	s.router.HandleFunc("/api/v1/links/get", s.getLink).Methods(http.MethodGet)
}

func (s *Service) createLink(w http.ResponseWriter, r *http.Request) {
	var link string
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.mu.Lock()
	s.links = append(s.links, link)
	json.NewEncoder(w).Encode(domain + encode(uint64(len(s.links))))
	s.mu.Unlock()
}

func (s *Service) getLink(w http.ResponseWriter, r *http.Request) {
	var link string
	err := json.NewDecoder(r.Body).Decode(&link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ind, err := decode(strings.TrimPrefix(link, domain))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.links[ind-1])
}

func encode(number uint64) string {
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(11)

	for ; number > 0; number = number / length {
		encodedBuilder.WriteByte(alphabet[(number % length)])
	}

	return encodedBuilder.String()
}

func decode(encoded string) (uint64, error) {
	var number uint64

	for i, symbol := range encoded {
		alphabeticPosition := strings.IndexRune(alphabet, symbol)

		if alphabeticPosition == -1 {
			return uint64(alphabeticPosition), errors.New("invalid character: " + string(symbol))
		}
		number += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return number, nil
}
