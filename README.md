# Go PostgreSQL SQL Builder example
PostgreSQL, SQL query builder and migrations usage example in Go

## Running Locally
* Run Postgres locally using Docker Compose: `make run-postgres`
* Install golang-migrate: `brew install golang-migrate`
* Apply migrations: `make migrate`
* Generate SQL Builder and Data Model types: `make jet-gen`
* Run the application: `go run main. go`
* Alternatively, run the app and PostgreSQL using Docker Compose: `make build && make run`

## Technologies Used
* [Go](https://go.dev/)
* [lib/pq](https://github.com/lib/pq)
* [go-jet](https://github.com/go-jet/jet/)
* [golang-migrate](https://github.com/golang-migrate/migrate)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)
