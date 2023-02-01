package controller

import (
	"forum/internal/entity"
	"forum/utils"
	"net/http"
)

type UserResponse struct {
	User         entity.User
	Posts        []entity.Post
	PostsLike    []entity.Post
	Postsdislike []entity.Post
}

func (h *Handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}

	// ts, err := template.ParseFiles(pageUser)
	// if err != nil {
	// 	h.InternalServerError500(w, r, err)
	// 	return
	// }

	var postsLike []entity.Post
	var postsDislike []entity.Post

	session, user, _ := h.getSessionFromCookie(w, r)

	// get posts by userId...
	posts, err := h.usecases.GetPostsByUserId(session.UserId)
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}

	// get likes by userId...
	likes, err := h.usecases.GetLikesByUserId(session.UserId)
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}

	for i := range likes {
		if likes[i].PostId != 0 {
			postLike, err := h.usecases.GetPostById(likes[i].PostId)
			if err != nil {
				h.InternalServerError500(w, r, err)
				return
			}
			postsLike = append(postsLike, postLike)
		}
	}

	// get dislikes by userId...
	dislikes, err := h.usecases.GetDislikesByUserId(session.UserId)
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}

	for i := range dislikes {
		if dislikes[i].PostId != 0 {
			postDislike, err := h.usecases.GetPostById(dislikes[i].PostId)
			if err != nil {
				h.InternalServerError500(w, r, err)
				return
			}
			postsDislike = append(postsDislike, postDislike)
		}
	}

	userOutput := UserResponse{
		User:         user,
		Posts:        posts,
		PostsLike:    postsLike,
		Postsdislike: postsDislike,
	}

	// err = ts.Execute(w, userOutput)
	// if err != nil {
	// 	h.InternalServerError500(w, r, err)
	// 	return
	// }

	if err := h.tmpl.ExecuteTemplate(w, "user.html", userOutput); err != nil {
		h.InternalServerError500(w, r, err)
		return
	}
}
