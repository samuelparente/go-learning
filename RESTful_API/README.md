# Go RESTful API Example

This is a simple RESTful API example for managing books in Go. It demonstrates CRUD (Create, Read, Update, Delete) operations on resources using Go's `net/http` package.

## Description

The API allows users to perform CRUD operations on books. It provides endpoints for creating, retrieving, updating, and deleting books. Books are represented as JSON objects with fields for ID, title, and author.

## Usage

API Endpoints
The API exposes the following endpoints:

Create a Book: POST /books/create
Get All Books: GET /books
Get a Book by ID: GET /books/get?id=<book_id>
Update a Book by ID: PUT /books/update?id=<book_id>
Delete a Book by ID: DELETE /books/delete?id=<book_id>

## Example Usage
Here are example curl commands to interact with the API:

### Create a Book:

```url -X POST -d '{"id": 1, "title": "Go Programming", "author": "John Doe"}' http://localhost:8080/books/create

### Get All Books:

```curl http://localhost:8080/books

### Get a Book by ID:

```curl http://localhost:8080/books/get?id=1

### Update a Book by ID:

```curl -X PUT -d '{"title": "Updated Title", "author": "Jane Smith"}' http://localhost:8080/books/update?id=1

### Delete a Book by ID:

```curl -X DELETE http://localhost:8080/books/delete?id=1