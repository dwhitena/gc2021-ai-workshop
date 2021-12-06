# Example 2 - Python (Flask) model server

This example illusrates how an AI model can be served using a custom Pythons server written in Flask. To run the example:

In Docker:

```
$ docker run -it -p 8080:8080 dwhitena/gc21-pyqa
```

Locally:
- Make sure Python 3 and `pip` are installed.
- Run:

    ```
    $ pip install -r requirements.txt
    $ python answer.py 
    ```

You can then interact with the server via cURL:

```
$ curl --location --request POST 'localhost:8080/answer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "context": "Go is a programming language designed at Google in 2007.",
    "question": "When was Go designed?"
}'
```
