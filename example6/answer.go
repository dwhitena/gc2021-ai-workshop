package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/nlpodyssey/spago/pkg/nlp/transformers/bert"
)

// Define pre-trained Hugging Face model.
var modelDir = "models/deepset/bert-base-cased-squad2"
var modelURL = "https://ai-classroom.nyc3.digitaloceanspaces.com/models.zip"
var model *bert.Model

func init() {

	// Check if the pre-trained model is available.
	if _, err := os.Stat(modelDir); os.IsNotExist(err) {
		if err := retrieveModel(modelURL); err != nil {
			log.Fatal(err)
		}
	}

	// Load the pre-trained model.
	modelLoad, err := bert.LoadModel(modelDir)
	if err != nil {
		log.Fatal(err)
	}
	model = modelLoad

	// Test the model after loading it in.
	question := "What is Go?"
	context := "Go is a programming language designed at Google in 2007."
	result := model.Answer(question, context)
	if result != nil && result[0].Confidence < 0.5 {
		log.Fatal("Couldn't run Q&A inference")
	}
	log.Println("Model ready for inference 🎉")
}

// retrieveModel calls a shell command wget to download and unzips
// the models file.
func retrieveModel(url string) error {

	// run the shell command `wget URL -O filepath`.
	cmd := exec.Command("wget", url, "-O", "models.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// unzip the models.zip file.
	cmd = exec.Command("unzip", "models.zip")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func main() {

	// ListenAndServe starts an HTTP server with a given address and
	// handler defined in NewRouter.
	log.Println("Starting service on port 8080")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
	defer model.Close()
}
