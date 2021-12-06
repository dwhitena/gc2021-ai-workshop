# Example 6 - AI in Go

This example illustrates how to import and call a Hugging Face model directly from Go. To run the application:

Locally:

1. Build the application

    ```
    $ go build
    ```

2. Run the application

    ```
    $ ./example6
    ```

In Docker:

```
$ docker run -it -p 8080:8080 dwhitena/gc21-goqa
```

Test the application:


```
$ go test -v
```

Call the API with cURL:


```
$ curl --location --request POST 'localhost:8080/answer' \
--header 'Content-Type: application/json' \
--data-raw '{
    "context": "Go is a programming language designed at Google in 2007.",
    "question": "When was Go designed?"
}'
```
