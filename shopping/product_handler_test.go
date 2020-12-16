package shopping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/matryer/is"
)

func TestUnit_Create(t *testing.T) {
	is := is.New(t)
	ph := createTestProductHandler()

	r := httptest.NewRequest("GET", "/any-url", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})

	w := httptest.NewRecorder()

	// see below:
	// this is just testing the handler func QueryByID
	// doesn't test whether the url route is working
	// TODO: try placing any url instead of /products and this will work
	// N.B. The `id` has to be passed as it's checked inside the func scope
	ph.QueryByID(w, r)

	is.Equal(w.Result().StatusCode, http.StatusOK)

}

func TestInt_Create(t *testing.T) {
	is := is.New(t)
	ph := createTestProductHandler()

	r := httptest.NewRequest("GET", "/products", nil)
	r = mux.SetURLVars(r, map[string]string{"id": "1"})

	w := httptest.NewRecorder()

	// see the diff below:
	// ph.QueryByID is passed into http.HandlerFunc()
	// this is a way of testing the whole stack using the router
	// e.g. integration and e2e tests (as seen in e2e_test.go)
	handler := http.HandlerFunc(ph.QueryByID)
	handler.ServeHTTP(w, r)

	is.Equal(w.Result().StatusCode, http.StatusOK)

}

func createTestProductHandler() *ProductHandler {
	pm := make(map[string]*Product)
	pm["1"] = &Product{Id: "1", Product: "abc", Price: "£1.00"}
	pm["2"] = &Product{Id: "2", Product: "def", Price: "£2.00"}

	mdb := NewDB(pm)
	ps := NewProductService(mdb)
	ph := NewProductHandler(ps)

	return ph
}
