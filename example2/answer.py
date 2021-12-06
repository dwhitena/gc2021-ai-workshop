import os
import time

from transformers.pipelines import pipeline
from flask import Flask
from flask import request, jsonify


#-----------------#
#   FLASK APP     #
#-----------------#

def create_app(hg_comp):

    # Create my flask app
    app = Flask(__name__)

    # Define a handler for the / health check.
    @app.route("/")
    def hello_world():
        return "<p>The question answering API is healthy!</p>"


    # Define a handler for the /answer path, which
    # processes a JSON payload with a question and
    # context and returns an answer using a Hugging
    # Face model.
    @app.route("/answer", methods=['POST'])
    def answer():

        # Get the request body data
        data = request.json

        # Answer the question
        answer = hg_comp({'question': data['question'], 
            'context': data['context']})

        # Create the response body.
        out = {
            "question": data['question'],
            "context": data['context'],
            "answers": [{
                "text": answer['answer'],
                "confidence": answer['score']
            }]
        }
        print(out)

        return jsonify(out)

    return app


#--------------#
#  FUNCTIONS   #
#--------------#

# Run main by default if running "python answer.py"
if __name__ == '__main__':

    # Initialize our default model.
    hg_comp = pipeline('question-answering', 
                    model="distilbert-base-uncased-distilled-squad", 
                    tokenizer="distilbert-base-uncased-distilled-squad")

    # Create our Flask App
    app = create_app(hg_comp)

    # Run our Flask app and start listening for requests!
    app.run(host='0.0.0.0', port=int(os.environ.get("PORT", 8080)), threaded=True)
