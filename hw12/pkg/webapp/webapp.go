package webapp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"homeworks/hw12/pkg/index"
	"net/http"
)

type Service struct {
	ind *index.Index
}

func New(ind *index.Index, r *mux.Router) Service {
	s := Service{ind: ind}
	s.endpoints(r)
	return s
}

func (s *Service) endpoints(r *mux.Router) {
	r.HandleFunc("/index", s.indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/docs", s.docsHandler).Methods(http.MethodGet)
}

func (s *Service) indexHandler(w http.ResponseWriter, _ *http.Request) {
	writeIndData(w, s.ind.Data)
}

func (s *Service) docsHandler(w http.ResponseWriter, _ *http.Request) {
	writeIndData(w, s.ind.Docs)
}

func writeIndData(w http.ResponseWriter, indData any) {
	data, err := json.Marshal(indData)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}
}
