# A Restful Go API
This is a simple example of a Go RESTful API

## How to run the file

To clean up unused dependencies and ensure that the `go.mod` and `go.sum` files are up to date, run the following command:

```bash
go mod tidy
```

To run the application, use the following command:

```bash
go run .
```

## How to test the API

The server is running at ``http://localhost:9090``

Open Postman and select the GET method to get the data.

Write this route to GET all data:

```bash
localhost:9090/todos
```
