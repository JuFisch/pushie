package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	// Create a request to pass to the handler.
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies HomeHandler's http.ResponseWriter argument?) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	// Use handler which takes as arguments (w Http.ResponseWriter, r *http.Request)

	handler.ServeHTTP(rr, req)

	// Check the return value of the handler is what we expect.
	status := rr.Code

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %s", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %s", rr.Body.String(), expected)
	}

}
