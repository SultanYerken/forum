package usecase

import (
	"forum/internal/entity"
	"forum/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	SignIn(username, password string) (string, error)
	ParseToken(token string) (entity.Session, error)
	LogOutUser(token string) error
}

type TodoPost interface {
	GetUser(userId int) (entity.User, error)
	CreatePost(user entity.User, post, category string) (int, error)
	DeletePost(postId int) error
	GetAllPost() ([]entity.Post, error)
	GetPostById(postId int) (entity.Post, error)
	GetCategoryId(categoryName string) (int, error)
	WriteInPostCategory(postId, categoryId int) (int, error)
	GetPostsIdByCategory(categoriesId []int) ([]int, error)
	GetPostsByUserId(userId int) ([]entity.Post, error)

	WriteLikeDislikeCountInPost(post entity.Post, postId int) (entity.Post, error)
}

type TodoComment interface {
	CreateComment(user entity.User, postId int, comment string) (int, error)
	GetAllComments(postId int) ([]entity.Comment, error)
	GetCommentById(commentId int) (entity.Comment, error)

	WriteLikeDislikeCountInComment(comment entity.Comment, commentId int) (entity.Comment, error)
}

type TodoLikeDislike interface {
	Like(userId, postId, commentId int) (int, error)
	Dislike(userId, postId, commentId int) (int, error)
	GetLikesByUserId(userId int) ([]entity.LikeDislike, error)
	GetDislikesByUserId(userId int) ([]entity.LikeDislike, error)
}

type UseCase struct {
	Authorization
	TodoPost
	TodoComment
	TodoLikeDislike
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		Authorization:   NewAuthUseCase(repos.Authorization),
		TodoPost:        NewTodoPostUseCase(repos.TodoPost),
		TodoComment:     NewTodoCommentUseCase(repos.TodoComment),
		TodoLikeDislike: NewTodoLikeDislikeUseCase(repos.TodoLikeDislike),
	}
}
