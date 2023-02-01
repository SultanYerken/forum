package usecase

import (
	"forum/internal/entity"
	"forum/internal/repository"
)

type TodoLikeDislikeUseCase struct {
	repo repository.TodoLikeDislike
}

func NewTodoLikeDislikeUseCase(repo repository.TodoLikeDislike) *TodoLikeDislikeUseCase {
	return &TodoLikeDislikeUseCase{repo: repo}
}

func (u *TodoLikeDislikeUseCase) Like(userId, postId, commentId int) (int, error) {
	res, err := u.repo.Checklike(userId, postId, commentId)
	if err != nil {
		return 0, err
	}
	if res == true {
		// log.Println("like deleted")
		return 0, nil
	}

	res, err = u.repo.Checkdislike(userId, postId, commentId)
	if err != nil {
		return 0, err
	}
	if res == true {
		// log.Println("dislike deleted")
	}

	return u.repo.LikeInDB(userId, postId, commentId)
}

func (u *TodoLikeDislikeUseCase) Dislike(userId, postId, commentId int) (int, error) {
	res, err := u.repo.Checkdislike(userId, postId, commentId)
	if err != nil {
		return 0, err
	}
	if res == true {
		// fmt.Println("dislike deleted")
		return 0, nil
	}

	res, err = u.repo.Checklike(userId, postId, commentId)
	if err != nil {
		return 0, err
	}
	if res == true {
		// fmt.Println("like deleted")
	}
	return u.repo.DislikeInDB(userId, postId, commentId)
}

func (u *TodoLikeDislikeUseCase) GetLikesByUserId(userId int) ([]entity.LikeDislike, error) {
	return u.repo.GetLikesByUserId(userId)
}

func (u *TodoLikeDislikeUseCase) GetDislikesByUserId(userId int) ([]entity.LikeDislike, error) {
	return u.repo.GetDislikesByUserId(userId)
}
