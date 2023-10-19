# Golang rest api 💻

## Task 📝

1. Create a simple Golang application that will serve a RESTful API with a resource (books). The endpoints should support Create, Read, Update and Delete operations.
2. The data should be persisted in a database like SQLite, MySQL, PostgreSQL, etc.
3. Add tests for you REST API.
4. Create a dockerfile for your application.

## Tips 🧞

### Application example

_Endpoints:_

```sh
# Get all books
GET /books

# Get a specific book
GET /books/:id

# Create a book
POST /books

# Update a book
PUT /books/:id

# Delete a book
DELETE /books/:id
```

_Model:_

```json
{
    "id": 1,
    "title": "The Alchemist",
    "author": "Paulo Coelho",
    "published": 1988
}
```

### Database

This package can be used to connect to a local SQLite database:

-   https://github.com/mattn/go-sqlite3

If you want an ORM (Object Relational Mapping) library, you can use GORM:

-   https://gorm.io/docs/connecting_to_the_database.html#SQLite

### REST API

For creating a HTTP server, you can use echo:

-   https://github.com/labstack/echo

### REST API testing

Documentation for testing echo endpoints:

-   https://echo.labstack.com/guide/testing/
