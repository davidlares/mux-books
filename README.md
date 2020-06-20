# Mux-books

This repo contains a simple API built with` Golang's Mux HTTP multiplexer` and a simple cloud instance of `PostgreSQL` with `ElephantSQL` service.

This API performs a REST API's CRUD operations for a "books" resource. the "GET", "PUT" and "DELETE" methods expects an `ID` parameter via URL query string.

The project structure is separated by its concerns, it uses a certain kind of MVC pattern for better scalability, for it was never intended for expected growth, it is a workshop project for learning Golang's HTTP handling.

## Usage

Before installing dependencies, you should have created a `.env` file key the key value of `ELEPHANTSQL_URL` key for the `PostgreSQL connection string`

Once the dependencies are installed, just:

`go run main.go`

And that's it.

## Dependencies

There's a lot of them:

1. `gorilla` as a MUX router
2. `net/http` as a request handler
3. `encoding/json` for JSON encoding,
4. `reflect` for type checking
5. `strconv` for string-int conversion

## Credits

 - [David E Lares](https://twitter.com/davidlares3)

## License

 - [MIT](https://opensource.org/licenses/MIT)
