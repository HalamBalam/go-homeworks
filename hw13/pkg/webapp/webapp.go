package webapp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"homeworks/hw13/pkg/api"
	"homeworks/hw13/pkg/index"
	"net/http"
)

type Service struct {
	ind *index.Index
	api *api.Service
}

func New(ind *index.Index, r *mux.Router) Service {
	s := Service{ind: ind}
	s.api = api.New(r, ind)
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
