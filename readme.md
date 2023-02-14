# Go learning by doing

## Project 1

[HTTP-server](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server)

Implementing a web server following the TDD approach with functionality where users can track how many games players have won.

* GET `/players/{name}` should return a number indicating the total number of wins
* POST `/players/{name}` should record a win for that name, incrementing for every subsequent POST

### Lesson 1

Date: 2023-01-30

**Summary:** Reading the server specifications

### Lesson 2

Continue from [write the test first](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#write-the-test-first)

* `go.mod` file is needed for golang to understand that the current location contains a golang project
* `net/http` has two methods `HandlerFunc` and `Handler` the difference between those two is that the fist one takes 
a function to handle received request and the second one takes the `Handler` interface implementation. When request 
handler doesn't need state then it's easier to use the `HandlerFunc` with a function to not overcomplicate the logic 
by implementing the only function `ServeHTTP` from the `Handler` interface 

### Lesson 3

Continue from [`http.HandlerFunc`](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#http.handlerfunc)
Open question: How is it possible for a function to be a receiver? How it works exactly?
