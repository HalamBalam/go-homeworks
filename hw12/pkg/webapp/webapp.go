package webapp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"homeworks/hw12/pkg/index"
	"net/http"
)

var Ind *index.Index

func Start() error {
	r := mux.NewRouter()
	endpoints(r)

	return http.ListenAndServe(":8080", r)
}

func endpoints(r *mux.Router) {
	r.HandleFunc("/index", indexHandler).Methods(http.MethodGet)
	r.HandleFunc("/docs", docsHandler).Methods(http.MethodGet)
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	writeIndData(w, Ind.Data)
}

func docsHandler(w http.ResponseWriter, _ *http.Request) {
	writeIndData(w, Ind.Docs)
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
