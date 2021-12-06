package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Index is the handler for the root URL.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the GopherCon 2021 Q&A API!\n")
}

// AnswerDetail stores individual answer details.
type AnswerDetail struct {
	Text       string  `json:"text"`
	Confidence float32 `json:"confidence"`
}

// Answer is used to pass predicted answers back to a user.
type Answer struct {
	Context  string         `json:"context"`
	Question string         `json:"question"`
	Answers  []AnswerDetail `json:"answers"`
}

// getInference is used to retrieve an inference from the
// Python-based inference server.
func getInference(context, question string) (Answer, error) {

	// Prepare the payload.
	payload, err := json.Marshal(map[string]string{
		"context":  context,
		"question": question,
	})
	if err != nil {
		return Answer{}, err
	}

	// Prepare the client and request.
	client := &http.Client{}
	req, err := http.NewRequest("POST", os.Getenv("INFER_URL"), bytes.NewBuffer(payload))
	if err != nil {
		return Answer{}, err
	}
	req.Header.Add("Content-Type", "application/json")

	// Make the request.
	res, err := client.Do(req)
	if err != nil {
		return Answer{}, err
	}
	defer res.Body.Close()

	// Parse the body of the response
	var answer Answer
	if err = json.NewDecoder(res.Body).Decode(&answer); err != nil {
		return Answer{}, err
	}
	return answer, nil
}

// LangDetect processes one or more strings, provided by a user and
// returns detected languages corresponding to each string.
func AnswerQuestion(w http.ResponseWriter, r *http.Request) {

	// Create an Answer value to parse the user input,
	// and eventually hold the returned answer.
	var inAnswer Answer

	// Parse the body of the request.
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Printf("%s: %s", "ERROR could not parse request body", err.Error())
	}
	if err := r.Body.Close(); err != nil {
		log.Printf("%s: %s", "ERROR could not close request body", err.Error())
	}
	if err := json.Unmarshal(body, &inAnswer); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("%s: %s", "ERROR could not unmarshal input JSON", err.Error())
		}
		return
	}

	// Get the answer.
	outAnswer, err := getInference(inAnswer.Context, inAnswer.Question)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("%s: %s", "ERROR could not encode the output JSON", err.Error())
		}
		return
	}

	// Send back the answer!
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(outAnswer); err != nil {
		log.Printf("%s: %s", "ERROR could not encode the output JSON", err.Error())
	}
}

// NotFound responds to problematic requests.
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	notFound := map[string]string{
		"Method":     r.Method,
		"URL":        r.URL.String(),
		"RequestURI": r.RequestURI,
	}
	if err := json.NewEncoder(w).Encode(notFound); err != nil {
		log.Printf("%s: %s", "ERROR could not marshal not found message to JSON", err.Error())
	}
}
