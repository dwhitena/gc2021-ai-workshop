# Example 4 - Hosted Inference Service + Go integration

This example illustrates:

1. Configuring a hosted inference service (Hugging Face)
2. Calling a hosted inference service (Hugging Face) from Go

To run the example:

1. Create an account on [Hugging Face](https://huggingface.co/)
2. Go to your account settings
3. Create a new access token, copy the token
3. Build the Go application:

    ```
    $ go build
    ```

3. Run the Go application, replacing `<access token>` with your access token:

    ```
    $ INFER_URL=https://api-inference.huggingface.co/models/distilbert-base-cased-distilled-squad \
      INFER_KEY=<access token> \
      ./example4
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
