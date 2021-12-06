# GopherCon 2021 "Production AI with Go" Workshop

Machine learning and artificial intelligence (ML/AI) applications are increasingly colliding with the systems that we maintain. These applications have a bunch of quirks that make them rather hard to scale, debug, and integrate. ML/AI applications are often written in Python, rely on a variety of large serialized assets (like model binaries), use specialized hardware (like GPUs), and exhibit behavior that is hard to predict. In this workshop, you will learn how to successfully integrate Go applications with common ML/AI components like TensorFlow and PyTorch to create an AI-enabled API for automated question answering. Along the way, you will learn how to appropriately test ML/AI components, collaborate with data science teams, and navigate the weird toolkit used by AI practitioners. 

In addition to a general understanding of common ML/AI components and their integration with Go applications, you will learn about testing and monitoring methodologies for AI models, common AI inference server infrastructure, and best practices for running AI in production. You will leave the workshop having built a Go API that calls into a state-of-the-art AI model to enable automated question answering.

*Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class via slides, discussion, etc.*

## Agenda

1. What is AI?
2. What does it mean to run AI in production?
    1. [Example 1](example1) - Python notebook
3. How can AI models be integrated into Go applications?
4. Introduction to our example for the workshop
5. Hands-on examples/ coding:
    1. [Example 2](example2) - Python model server (Flask)
    2. [Example 3](example3) - Python + Go, Testing
    3. [Example 4](example4) - Hosted Inference Service (Hugging Face Infernce API) + Go
    4. [Example 5](example5) - Importing Hugging Face models into Go
    5. [Example 6](example6) - AI in Go with spaGO
6. Conclusions/ Q&A

## Prerequisites

You do not need to be a Go or AI expert for this training and beginners are highly encouraged to attend. You do need some exposure to Go and must have at least gone through the [Go Tour](https://tour.golang.org/).

Recommended Preparation: 
- Install and configure an editor for Go.
- Have a functioning Go environment installed with Go 1.16 or later.
- Sign up for a GitHub account, if you do not already have one.
- Computers should be capable of modern software development, such as access to install and run binaries, install a code editor, etc. 
- Though not completely necessary, having Docker installed, will also be useful.


