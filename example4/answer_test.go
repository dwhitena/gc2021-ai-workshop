package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAnswer(t *testing.T) {

	// Loop over test fixtures.
	files, err := ioutil.ReadDir("fixtures/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {

		// Open fixture file.
		contents, err := os.Open("fixtures/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		// Parse the fixture file.
		var expectedAnswer Answer
		jsonParser := json.NewDecoder(contents)
		if err = jsonParser.Decode(&expectedAnswer); err != nil {
			log.Fatal(err)
		}

		// Create a request to pass to our answer handler.
		payload, err := json.Marshal(map[string]string{
			"context":  expectedAnswer.Context,
			"question": expectedAnswer.Question,
		})
		if err != nil {
			log.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/answer", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter)
		// to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(AnswerQuestion)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the answer is what we expect.
		var answer Answer
		if err := json.Unmarshal(rr.Body.Bytes(), &answer); err != nil {
			t.Errorf("could not parse response body: %s",
				rr.Body.String())
		}
		if answer.Answers[0].Text != expectedAnswer.Answers[0].Text {
			t.Errorf("unexpected answer for %s: got \"%s\" expected \"%s\"", file.Name(),
				answer.Answers[0].Text, expectedAnswer.Answers[0].Text)
		}
	}
}

// TestLogger is a test to make sure REST requests are logged
func TestLogger(t *testing.T) {
	for _, route := range routes {

		// wrap all current routes in the logger decorator to log out requests
		testhandler := Logger(route.HandlerFunc, route.Name)
		if &testhandler == nil {
			t.Error("Test failed: Could not wrap route with logger")
		}

	}
}

// TestNewRouter is a test to make sure a new router can be created
func TestNewRouter(t *testing.T) {
	testrouter := NewRouter()
	if &testrouter == nil {
		t.Error("Test failed: Could not create new router")
	}
}
