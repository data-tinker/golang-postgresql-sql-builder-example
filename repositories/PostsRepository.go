package repositories

import (
	"context"
	"database/sql"
	"golang-postgresql-sql-builder-example/.gen/blog/public/model"
	"golang-postgresql-sql-builder-example/.gen/blog/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

type PostsRepository struct {
	Db *sql.DB
}

func (r *PostsRepository) GetPostByID(id int32) (*model.Posts, error) {
	query := SELECT(table.Posts.AllColumns).
		FROM(table.Posts).
		WHERE(table.Posts.ID.EQ(Int32(id)))

	post := &model.Posts{}
	err := query.Query(r.Db, post)

	return post, err
}

func (r *PostsRepository) CreatePost(userId int32, title string, content string) (*model.Posts, error) {
	insertStmt := table.Posts.INSERT(
		table.Posts.UserID,
		table.Posts.Title,
		table.Posts.Content,
	).VALUES(
		userId,
		title,
		content,
	).RETURNING(table.Posts.AllColumns)

	dest := &model.Posts{}

	err := insertStmt.Query(r.Db, dest)

	return dest, err
}

func (r *PostsRepository) UpdatePosts(ctx context.Context, posts []*model.Posts) error {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, post := range posts {
		updateStmt := table.Posts.
			UPDATE(table.Posts.Title, table.Posts.Content).
			SET(post.Title, post.Content).
			WHERE(table.Posts.ID.EQ(Int32(post.ID)))

		_, err := updateStmt.Exec(tx)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
