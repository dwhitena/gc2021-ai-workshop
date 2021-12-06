# Example 5 - Importing Hugging Face Models into Go

There are a variety of Go packages that allow you to import AI models (trained in TensorFlow, PyTorch, or similar) into Go. By way of example, we will use [github.com/nlpodyssey/spago](https://github.com/nlpodyssey/spago), which integrates with certain Hugging Face models. The developers of `spaGO` have created [a utility](https://github.com/nlpodyssey/spago/tree/main/cmd/huggingfaceimporter) which will allow you to convert a Hugging Face model into a `spaGO` model. We will use this utility here to convert a Hugging Face Q&A model into `spaGO` format. 

*Note - The compiled binary for this utility (included here) is compiled for Linux systems. Compile the [spaGO source](https://github.com/nlpodyssey/spago/tree/main/cmd/huggingfaceimporter) if you want to use it on another system.*

To run the example:

```
$ ./get-model.sh
```
