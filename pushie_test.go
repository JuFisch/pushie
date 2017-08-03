package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//"github.com/stretchr/testify/assert"

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

func TestPublishHandler(t *testing.T) {

	fmt.Println("TestPublishHandler says: Running TestRoutes")

	// Set up body with message.
	message := Message{
		Channel: "bunno",
		Data: map[string]interface{}{
			"message": map[string]string{"text": "oh hai", "username": "jufisch"},
		},
	}

	// Encode message to messagebytes ([]byte).
	mbytes, _ := json.Marshal(message)
	fmt.Println("TestPublishHandler says: The encoded JSON for message is", mbytes)
	//message_bytes := []byte(`{"blah": "blah"}`)
	mbuffer := bytes.NewBuffer(mbytes)

	// Set up the request with the path and the message.
	req, err := http.NewRequest("POST", "/publish", mbuffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TestPublishHandler says: The final request sent is ", req)
	// req, _ := http.NewRequest(method, urlStr, body)

	// Set up expected response.
	expected_response := []byte(`{"sockets_count":2}`)

	// response writer
	// call our function
	// test output
	var res http.ResponseWriter

	PublishHandler(res, req)
	// Assert that res got a message.
	fmt.Printf("TestPublishHandler says: Assert we got '%s'.\n", expected_response)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PublishHandler)

	handler.ServeHTTP(rr, req)

	//  assert.Equal(t, a, b, "The two words should be the same.")

}
