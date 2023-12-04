run-postgres:
	docker-compose up postgres

migrate:
	migrate -path=migrations/ -database "postgres://postgres:password@localhost:5432/blog?sslmode=disable" up

jet-gen:
	jet -dsn=postgres://postgres:password@localhost:5432/blog?sslmode=disable -schema=public -path=./.gen

build:
	docker build -t golang-postgresql-sql-builder-example .

run:
	docker-compose up
