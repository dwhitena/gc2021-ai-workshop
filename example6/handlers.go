package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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
// spaGO implementation of the model.
func getInference(context, question string) Answer {

	// Initialize an Answer value.
	var answer Answer
	answer.Question = question
	answer.Context = context

	// Predict an answer to the input question.
	result := model.Answer(question, context)
	answers := []AnswerDetail{}
	for _, answer := range result {
		answers = append(answers, AnswerDetail{
			Text:       answer.Text,
			Confidence: answer.Confidence,
		})
	}
	answer.Answers = answers
	return answer
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Printf("%s: %s", "ERROR could not unmarshal input JSON", err.Error())
		}
		return
	}

	// Predict an answer to the input question.
	outAnswer := getInference(inAnswer.Context, inAnswer.Question)

	// Return the result.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(outAnswer); err != nil {
		log.Printf("%s: %s", "ERROR could not unmarshal input JSON", err.Error())
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
