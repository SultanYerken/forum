package usecase

import (
	"forum/internal/entity"
	"forum/internal/repository"
)

type TodoCommentUseCase struct {
	repo repository.TodoComment
}

func NewTodoCommentUseCase(repo repository.TodoComment) *TodoCommentUseCase {
	return &TodoCommentUseCase{repo: repo}
}

func (u *TodoCommentUseCase) CreateComment(user entity.User, postId int, comment string) (int, error) {
	return u.repo.CreateComment(user, postId, comment)
}

func (u *TodoCommentUseCase) GetAllComments(postId int) ([]entity.Comment, error) {
	comments, err := u.repo.GetAllComments(postId)
	if err != nil {
		return comments, err
	}

	for i := range comments {
		comment, err := u.WriteLikeDislikeCountInComment(comments[i], comments[i].Id)
		if err != nil {
			return comments, err
		}
		comments[i] = comment
	}
	return comments, nil
}

func (u *TodoCommentUseCase) GetCommentById(commentId int) (entity.Comment, error) {
	return u.repo.GetCommentById(commentId)
}

func (u *TodoCommentUseCase) WriteLikeDislikeCountInComment(comment entity.Comment, commentId int) (entity.Comment, error) {
	likeComment, err := u.repo.GetLikesByCommentId(commentId)
	if err != nil {
		return comment, err
	}
	countlikes := len(likeComment)
	comment.LikeCount = countlikes

	dislikeComment, err := u.repo.GetDislikesByCommentId(commentId)
	if err != nil {
		return comment, err
	}
	countdislikes := len(dislikeComment)
	comment.DislikeCount = countdislikes

	return comment, nil
}
