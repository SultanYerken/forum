package repository

import (
	"database/sql"
	"forum/internal/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username string) (entity.User, error)
	CreateSession(userId int, token string) (int, string, error)
	GetSession(token string) (entity.Session, error)
	DeleteSession(token string) error
}

type TodoPost interface {
	GetUser(userId int) (entity.User, error)
	CreatePost(user entity.User, post, category string) (int, error)
	DeletePost(postId int) error
	GetAllPost() ([]entity.Post, error)
	GetPostById(postId int) (entity.Post, error)
	GetCategoryId(categoryName string) (int, error)
	WriteInPostCategory(postId, categoryId int) (int, error)
	GetPostIdByCategoryId(categoryId int) ([]int, error)
	GetPostsByUserId(userId int) ([]entity.Post, error)

	GetLikesByPostId(postId int) ([]entity.LikeDislike, error)
	GetDislikesByPostId(postId int) ([]entity.LikeDislike, error)
}

type TodoComment interface {
	CreateComment(user entity.User, postId int, comment string) (int, error)
	GetAllComments(postId int) ([]entity.Comment, error)
	GetCommentById(commentId int) (entity.Comment, error)

	GetLikesByCommentId(commentId int) ([]entity.LikeDislike, error)
	GetDislikesByCommentId(commentId int) ([]entity.LikeDislike, error)
}

type TodoLikeDislike interface {
	LikeInDB(userId, postId, commentId int) (int, error)
	DislikeInDB(userId, postId, commentId int) (int, error)
	Checklike(userId, postId, commentId int) (bool, error)
	Checkdislike(userId, postId, commentId int) (bool, error)
	GetLikesByUserId(userId int) ([]entity.LikeDislike, error)
	GetDislikesByUserId(userId int) ([]entity.LikeDislike, error)
}

type Repository struct {
	Authorization
	TodoPost
	TodoComment
	TodoLikeDislike
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthDataBase(db),
		TodoPost:        NewTodoPostDataBase(db),
		TodoComment:     NewTodoCommentDataBase(db),
		TodoLikeDislike: NewTodoLikeDislikeDataBase(db),
	}
}
