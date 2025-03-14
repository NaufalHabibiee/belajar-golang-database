package repository

import (
	"belajar-golang-database/entitiy"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entitiy.Comment) (entitiy.Comment, error)
	FindById(ctx context.Context, id int32) (entitiy.Comment, error)
	FindAll(ctx context.Context) ([]entitiy.Comment, error)
}
