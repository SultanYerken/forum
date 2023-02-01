package controller

import (
	"forum/utils"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}
	_, user, _ := h.getSessionFromCookie(w, r)

	postId, err := strconv.Atoi(r.FormValue("postId"))
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}

	// did not enter the Id of non-existent posts...
	post, err := h.usecases.GetPostById(postId)
	if err != nil {
		h.BadRequest400(w, r, utils.ErrPostNotExist)
		return
	}

	comment := r.FormValue("comment")

	if chekEmptyField([]string{comment}) {
		h.BadRequest400(w, r, utils.ErrEmptyFields)
		return
	}

	_, err = h.usecases.CreateComment(user, postId, comment)
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}

	log.Printf("User %s commented out %s post \n", user.Username, post.Username)

	http.Redirect(w, r, "/posts?id="+strconv.Itoa(postId), http.StatusSeeOther)
}
