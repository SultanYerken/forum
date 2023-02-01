package controller

import (
	"forum/utils"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) Like(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}

	session, user, _ := h.getSessionFromCookie(w, r)
	userId := session.UserId
	var postId, commentId int

	postIdfromForm := r.FormValue("postId")
	commentIdfromForm := r.FormValue("commentId")

	if postIdfromForm != "" {
		id, err := strconv.Atoi(postIdfromForm)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		postId = id

		// did not enter the Id of non-existent posts...
		post, err := h.usecases.GetPostById(postId)
		if err != nil {
			h.BadRequest400(w, r, utils.ErrPostNotExist)
			return
		}

		_, err = h.usecases.Like(userId, postId, 0)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		log.Printf("User %s liked (or delete like) %s post", user.Username, post.Username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if commentIdfromForm != "" {
		id, err := strconv.Atoi(commentIdfromForm)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		commentId = id

		// did not enter the Id of non-existent comments...
		comment, err := h.usecases.GetCommentById(commentId)
		if err != nil {
			h.BadRequest400(w, r, utils.ErrCommentNotExist)
			return
		}

		_, err = h.usecases.Like(userId, 0, commentId)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		log.Printf("User %s liked (or delete like) %s comment", user.Username, comment.Username)
		http.Redirect(w, r, "/posts?id="+strconv.Itoa(comment.PostId), http.StatusSeeOther)
	}
}

func (h *Handler) Dislike(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}

	session, user, _ := h.getSessionFromCookie(w, r)
	userId := session.UserId
	var postId, commentId int

	postIdfromForm := r.FormValue("postId")
	commentIdfromForm := r.FormValue("commentId")

	if postIdfromForm != "" {
		id, err := strconv.Atoi(postIdfromForm)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		postId = id

		// did not enter the Id of non-existent posts...
		post, err := h.usecases.GetPostById(postId)
		if err != nil {
			h.BadRequest400(w, r, utils.ErrPostNotExist)
			return
		}

		_, err = h.usecases.Dislike(userId, postId, 0)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		log.Printf("User %s disliked (or delete dislike) %s post", user.Username, post.Username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	if commentIdfromForm != "" {
		id, err := strconv.Atoi(commentIdfromForm)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		commentId = id

		// did not enter the Id of non-existent comments...
		comment, err := h.usecases.GetCommentById(commentId)
		if err != nil {
			h.BadRequest400(w, r, utils.ErrCommentNotExist)
			return
		}

		_, err = h.usecases.Dislike(userId, 0, commentId)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		log.Printf("User %s disliked (or delete dislike) %s comment", user.Username, comment.Username)
		http.Redirect(w, r, "/posts?id="+strconv.Itoa(comment.PostId), http.StatusSeeOther)
	}
}
