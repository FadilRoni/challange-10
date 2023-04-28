"# challange-10" 

# Book API GoLang

## How to use

```sh
  go run main.go
  ```

  or

```sh
  go run .
  ```

## Available Endpoints:

### Get All Books: ( GET Request )
* `localhost:8080/books`

### Get single Book ( GET request )
* `localhost:8080/books/:id`

### Add New Book: ( POST Request )
* `localhost:8080/books`

### Update Book: ( PUT Request )
* `localhost:8080/books/:id`

### Delete Book: ( Delete Request )
* `localhost:8080/books/:id`

## Running Test :

```sh
  go test -v ./service
  ```
