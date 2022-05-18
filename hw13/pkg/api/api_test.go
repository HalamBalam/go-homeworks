package api

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"homeworks/hw13/pkg/crawler"
	"homeworks/hw13/pkg/index"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var api *Service

func TestMain(m *testing.M) {
	var docs []crawler.Document
	docs = append(docs,
		crawler.Document{URL: "https://go.dev", Title: "The Go Programming Language"},
		crawler.Document{URL: "https://teachbase.ru/", Title: "Platform for learning courses"},
	)
	ind := index.New()
	ind.AddDocs(docs)
	r := mux.NewRouter()
	api = New(r, ind)

	m.Run()
}

func TestService_search(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/search/go", nil)

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	got := rr.Body.String()
	want := `["https://go.dev"]` + "\n"
	if got != want {
		t.Errorf("ответ неверен: получили %s, а хотели %s", got, want)
	}

	t.Log("Response: ", rr.Body)

	//=========================================

	req = httptest.NewRequest(http.MethodGet, "/search/python", nil)

	rr = httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	got = rr.Body.String()
	want = "[]\n"
	if got != want {
		t.Errorf("ответ неверен: получили %s, а хотели %s", got, want)
	}

	t.Log("Response: ", rr.Body)
}

func TestService_createDoc(t *testing.T) {
	l := len(api.ind.Docs)
	d := crawler.Document{URL: "https://www.google.com/", Title: "Поиск в Google"}
	payload, _ := json.Marshal(d)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/docs", bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	got := rr.Body.String()
	want := strconv.Itoa(l+1) + "\n"
	if !(got == want) {
		t.Errorf("ответ неверен: получили %s, а хотели %s", got, want)
	}

	got = api.ind.Docs[l].URL
	want = "https://www.google.com/"
	if !(got == want) {
		t.Errorf("URL нового документа неверен: получили %s, а хотели %s", got, want)
	}

	got = api.ind.Docs[l].Title
	want = "Поиск в Google"
	if !(got == want) {
		t.Errorf("Заголовок нового документа неверен: получили %s, а хотели %s", got, want)
	}

	t.Log("Response: ", rr.Body)
}

func TestService_updateDocFull(t *testing.T) {
	id := 1
	d := crawler.Document{URL: "https://www.google.com/"}
	payload, _ := json.Marshal(d)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/docs/"+strconv.Itoa(id), bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	got := api.ind.Docs[id-1].URL
	want := "https://www.google.com/"
	if !(got == want) {
		t.Errorf("URL документа неверен: получили %s, а хотели %s", got, want)
	}

	got = api.ind.Docs[id-1].Title
	want = ""
	if !(got == want) {
		t.Errorf("Заголовок документа неверен: получили %s, а хотели %s", got, want)
	}
}

func TestService_updateDocPart(t *testing.T) {
	id := 2
	d := crawler.Document{URL: "https://www.google.com/"}
	payload, _ := json.Marshal(d)
	req := httptest.NewRequest(http.MethodPatch, "/api/v1/docs/"+strconv.Itoa(id), bytes.NewBuffer(payload))

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	got := api.ind.Docs[id-1].URL
	want := "https://www.google.com/"
	if !(got == want) {
		t.Errorf("URL документа неверен: получили %s, а хотели %s", got, want)
	}

	got = api.ind.Docs[id-1].Title
	want = "Platform for learning courses"
	if !(got == want) {
		t.Errorf("Заголовок документа неверен: получили %s, а хотели %s", got, want)
	}
}

func TestService_deleteDoc(t *testing.T) {
	id := 1
	l := len(api.ind.Docs)
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/docs/"+strconv.Itoa(id), nil)

	rr := httptest.NewRecorder()

	api.router.ServeHTTP(rr, req)

	if !(rr.Code == http.StatusOK) {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	got := len(api.ind.Docs)
	want := l - 1
	if !(got == want) {
		t.Errorf("Количество документов в индексе неверно: получили %d, а хотели %d", got, want)
	}
}
