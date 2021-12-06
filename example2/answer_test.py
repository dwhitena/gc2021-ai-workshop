import os
import time

from transformers.pipelines import pipeline
import pytest
from answer import create_app

# initialize testing environment
hg_comp = pipeline('question-answering', 
                    model="distilbert-base-uncased-distilled-squad", 
                    tokenizer="distilbert-base-uncased-distilled-squad")

@pytest.fixture
def client():
    app = create_app(hg_comp)
    app.config["TESTING"] = True
    with app.test_client() as client:
        yield client

# Health check route test
def test_health(client):
    r = client.get("/")
    assert 200 == r.status_code

# Answer question route test
def test_answer(client):
    payload = {
        "question": "who did holly matthews play in waterloo rd?",
        "context": "She attended the British drama school East 15 in 2005, and left after winning a high-profile role in the BBC drama Waterloo Road, playing the bully Leigh-Ann Galloway.[6] Since that role, Matthews has continued to act in BBC's Doctors, playing Connie Whitfield; in ITV's The Bill playing drug addict Josie Clarke; and she was back in the BBC soap Doctors in 2009, playing Tansy Flack."
    }
    r = client.post("/answer", json=payload)
    assert 200 == r.status_code

