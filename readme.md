# Go learning by doing

## Project 1

[HTTP-server](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server)

Implementing a web server following the TDD approach with functionality where users can track how many games players have won.

* GET `/players/{name}` should return a number indicating the total number of wins
* POST `/players/{name}` should record a win for that name, incrementing for every subsequent POST