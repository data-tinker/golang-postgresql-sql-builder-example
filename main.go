package main

import (
	"context"
	"database/sql"
	"golang-postgresql-sql-builder-example/.gen/blog/public/model"
	"golang-postgresql-sql-builder-example/repositories"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func openDbConnection() *sql.DB {
	postgresUrl := os.Getenv("POSTGRES_URL")
	if postgresUrl == "" {
		postgresUrl = "postgres://postgres:password@localhost:5432/blog?sslmode=disable"
	}

	db, openErr := sql.Open("postgres", postgresUrl)
	if openErr != nil {
		log.Fatal("Error opening database: ", openErr)
	}
	return db
}

func main() {
	db := openDbConnection()
	defer db.Close()

	usersRepo := &repositories.UsersRepository{Db: db}
	user, _ := usersRepo.CreateUser("Alex", "asmartishin@gmail.com")
	log.Printf("New user: %v\n", user)

	postsRepo := &repositories.PostsRepository{Db: db}
	post1, _ := postsRepo.CreatePost(user.ID, "My new post 1", "Post 1")
	log.Printf("New post 1: %v\n", post1)
	post1.Title = "My updated post 1"

	post2, _ := postsRepo.CreatePost(user.ID, "My new post 2", "Post 2")
	log.Printf("New post 2: %v\n", post2)
	post2.Title = "My updated post 2"

	ctx := context.Background()
	postsToUpdate := []*model.Posts{post1, post2}
	_ = postsRepo.UpdatePosts(ctx, postsToUpdate)

	updatedPost1, err := postsRepo.GetPostByID(post1.ID)
	if err != nil {
		log.Fatal("Error getting post: ", err)
	}
	log.Printf("Update post 1: %v\n", updatedPost1)
}
