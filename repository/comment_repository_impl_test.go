package repository

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/entitiy"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertId(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment := entitiy.Comment{
		Email:   "repository@gmail.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 37) // berarti ingin mencari id 37
	if err != nil {
		panic(err)
	}
	
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

   comments, err := commentRepository.FindAll(context.Background())
   if err != nil {
	panic(err)
   }

   for _, comment := range comments {
	fmt.Println(comment)
   }
}