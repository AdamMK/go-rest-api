## Go REST

This is a project of REST API performing CRUD operations written in Golang.

## Installation and running

* Clone git repo
* Navigate to the project folder
* Resolve all external dependencies with

```bash
go mod tidy
```

* Run the application with
```bash
go run .
```

## Usage

Application exposes a number of endpoints. For better experience please use API testing tools e.g Postman  

* LIST all documents with GET method on `http://localhost:8282/documents`

* FIND single document with GET method `http://localhost:8282/documents/{id}`
    `id` is a unique UUID of each document passed as path parameter e.g `bd84ba76-fd80-43ae-b35c-0e2857d3b7aa`

* CREATE a document with POST method hitting `http://localhost:8282/documents` and passing a json in the body of the request
    example of the request body:
    ```json
    {
        "title": "Renting contract",
        "content": {
            "header": "Rent a tent",
            "data": "One person tent"
        },
        "signee": "Yogi Bear"
    }
    ```
* DELETE a document by using DELETE method on `http://localhost:8282/documents/{id}`
* UPDATE a document with PUT method `http://localhost:8282/documents/{id}` with path parameter of document to be changed along with the json body of the changes to make.

Examples of the data to play with can be found in the `data.json` file.

## Docker

This was my try to containarise the app. The docker container is created and application inside is working however for some reason the port of the container cannot be forwarded to local host. Any ideas are welcomed.