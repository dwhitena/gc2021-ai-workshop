# Example 3 - Python + Go integration

This example illustrates:

1. How to call a Python model server from Go
2. Test an AI-enabled application

To run the example:

1. Run the Python model server:

    ```
    $ docker run -it -p 8081:8080 dwhitena/gc21-pyqa
    ```

2. Build the Go application:

    ```
    $ go build
    ```

3. Run the Go application:

    ```
    $ INFER_URL=http://localhost:8081/answer ./example3
    ```

To test the API, run:

```
$ curl --location --request POST 'localhost:8080/answer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "context": "Go is a programming language designed at Google in 2007.",
    "question": "When was Go designed?"
}'
```
