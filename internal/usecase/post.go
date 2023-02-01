package usecase

import (
	"forum/internal/entity"
	"forum/internal/repository"
)

type TodoPostUseCase struct {
	repo repository.TodoPost
}

func NewTodoPostUseCase(repo repository.TodoPost) *TodoPostUseCase {
	return &TodoPostUseCase{repo: repo}
}

func (u *TodoPostUseCase) GetUser(userId int) (entity.User, error) {
	return u.repo.GetUser(userId)
}

func (u *TodoPostUseCase) CreatePost(user entity.User, post, category string) (int, error) {
	return u.repo.CreatePost(user, post, category)
}

func (u *TodoPostUseCase) DeletePost(postId int) error {
	return u.repo.DeletePost(postId)
}

func (u *TodoPostUseCase) GetAllPost() ([]entity.Post, error) {
	posts, err := u.repo.GetAllPost()
	if err != nil {
		return posts, err
	}

	for i := range posts {
		post, err := u.WriteLikeDislikeCountInPost(posts[i], posts[i].Id)
		if err != nil {
			return posts, err
		}
		posts[i] = post
	}

	return posts, nil
}

func (u *TodoPostUseCase) GetPostById(postId int) (entity.Post, error) {
	post, err := u.repo.GetPostById(postId)
	if err != nil {
		return post, err
	}

	post, err = u.WriteLikeDislikeCountInPost(post, postId)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (u *TodoPostUseCase) GetCategoryId(categoryName string) (int, error) {
	return u.repo.GetCategoryId(categoryName)
}

func (u *TodoPostUseCase) WriteInPostCategory(postId, categoryId int) (int, error) {
	return u.repo.WriteInPostCategory(postId, categoryId)
}

func (u *TodoPostUseCase) GetPostsIdByCategory(categoriesId []int) ([]int, error) {
	var allPostsId []int
	for i := range categoriesId {
		postsId, err := u.repo.GetPostIdByCategoryId(categoriesId[i])
		if err != nil {
			return nil, err
		}

		if postsId != nil {
			for i := range postsId {
				if chekRepeatPosts(allPostsId, postsId[i]) {
					continue
				}
				allPostsId = append(allPostsId, postsId[i])
			}
		}
	}
	return allPostsId, nil
}

func (u *TodoPostUseCase) GetPostsByUserId(userId int) ([]entity.Post, error) {
	return u.repo.GetPostsByUserId(userId)
}

func chekRepeatPosts(allPostsId []int, postId int) bool {
	for i := range allPostsId {
		if allPostsId[i] == postId {
			return true
		}
	}
	return false
}

func (u *TodoPostUseCase) WriteLikeDislikeCountInPost(post entity.Post, postId int) (entity.Post, error) {
	likesPost, err := u.repo.GetLikesByPostId(postId)
	if err != nil {
		return post, err
	}
	countlikes := len(likesPost)
	post.LikeCount = countlikes

	dislikesPost, err := u.repo.GetDislikesByPostId(postId)
	if err != nil {
		return post, err
	}
	countdislikes := len(dislikesPost)
	post.DislikeCount = countdislikes

	return post, nil
}
