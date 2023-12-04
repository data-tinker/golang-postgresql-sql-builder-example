package repositories

import (
	"database/sql"
	"golang-postgresql-sql-builder-example/.gen/blog/public/model"
	"golang-postgresql-sql-builder-example/.gen/blog/public/table"
)

type UsersRepository struct {
	Db *sql.DB
}

func (r *UsersRepository) CreateUser(username, email string) (*model.Users, error) {

	insertStmt := table.Users.
		INSERT(table.Users.Username, table.Users.Email).
		VALUES(username, email).
		RETURNING(table.Users.AllColumns)

	dest := &model.Users{}

	err := insertStmt.Query(r.Db, dest)

	return dest, err
}
