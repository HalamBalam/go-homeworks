package webapp

import (
	"github.com/gorilla/mux"
	"homeworks/hw12/pkg/crawler"
	"homeworks/hw12/pkg/index"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r *mux.Router

func TestMain(m *testing.M) {
	var docs []crawler.Document
	docs = append(docs,
		crawler.Document{URL: "https://go.dev", Title: "The Go Programming Language"},
		crawler.Document{URL: "https://teachbase.ru/", Title: "Platform for learning courses"},
	)

	r = mux.NewRouter()

	ind := index.New()
	ind.AddDocs(docs)
	New(ind, r)

	m.Run()
}

func Test_indexHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	req.Header.Add("content-type", "plain/text")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	want := `{"courses":[2],"for":[2],"go":[1],"language":[1],"learning":[2],"platform":[2],"programming":[1],"the":[1]}`
	if rr.Body.String() != want {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body, want)
	}

	t.Log("Response: ", rr.Body)
}

func Test_docsHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	req.Header.Add("content-type", "plain/text")

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	want := `[{"ID":1,"URL":"https://go.dev","Title":"The Go Programming Language","Body":""},{"ID":2,"URL":"https://teachbase.ru/","Title":"Platform for learning courses","Body":""}]`
	if rr.Body.String() != want {
		t.Errorf("ответ неверен: получили %s, а хотели %s", rr.Body, want)
	}

	t.Log("Response: ", rr.Body)
}
