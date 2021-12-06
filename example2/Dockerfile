FROM tensorflow/tensorflow:2.7.0

COPY requirements.txt . 

RUN pip install -r requirements.txt 

COPY answer.py /app/answer.py
COPY answer_test.py /app/answer_test.py
WORKDIR /app

CMD ["python3", "answer.py"]
