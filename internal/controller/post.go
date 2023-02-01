package controller

import (
	"forum/internal/entity"
	"forum/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ts, err := template.ParseFiles(pageCreatePost)
		// if err != nil {
		// 	h.InternalServerError500(w, r, err)
		// 	return
		// }
		// err = ts.Execute(w, nil)
		// if err != nil {
		// 	h.InternalServerError500(w, r, err)
		// 	return
		// }

		if err := h.tmpl.ExecuteTemplate(w, "create_post.html", nil); err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
	} else if r.Method == http.MethodPost {
		_, user, _ := h.getSessionFromCookie(w, r)

		// var postCategoryId int
		var categoryIds []int

		r.ParseForm()
		categories := r.Form["category"]
		post := r.FormValue("post")
		categorystr := strings.Join(categories, " ")

		if chekEmptyField([]string{post}) {
			h.BadRequest400(w, r, utils.ErrEmptyFields)
			return
		}

		postId, err := h.usecases.CreatePost(user, post, categorystr)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}

		if categories != nil {
			for i := range categories {
				categoryId, err := h.usecases.GetCategoryId(categories[i])
				if err != nil {
					err := h.usecases.DeletePost(postId)
					if err != nil {
						h.InternalServerError500(w, r, err)
						return
					}
					h.BadRequest400(w, r, utils.ErrWrongCategory)
					return
				}
				categoryIds = append(categoryIds, categoryId)
			}
			for i := range categoryIds {
				_, err := h.usecases.WriteInPostCategory(postId, categoryIds[i])
				if err != nil {
					h.InternalServerError500(w, r, err)
					return
				}

			}
		}

		log.Printf("post %d by %s created \n", postId, user.Username)

		http.Redirect(w, r, "/posts?id="+strconv.Itoa(postId), http.StatusSeeOther)

	} else {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}
}

type getPostByIdResponse struct {
	User     entity.User
	Post     entity.Post
	Comments []entity.Comment
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}

	// ts, err := template.ParseFiles(pagePost)
	// if err != nil {
	// 	h.InternalServerError500(w, r, err)
	// 	return
	// }

	_, user, err := h.getSessionFromCookie(w, r)
	if err != nil {
		if err == http.ErrNoCookie {
			// fmt.Println("User Unouthorized")
		}
	}

	postId, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		h.BadRequest400(w, r, utils.ErrBadRequest400)
		return
	}

	post, err := h.usecases.GetPostById(postId)
	if err != nil {
		h.NotFound404(w, r, utils.ErrPostNotExist)
		return
	}

	comments, err := h.usecases.GetAllComments(postId)
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}
	output := getPostByIdResponse{
		User:     user,
		Post:     post,
		Comments: comments,
	}

	// err = ts.Execute(w, output)
	// if err != nil {
	// 	h.InternalServerError500(w, r, err)
	// 	return
	// }

	if err := h.tmpl.ExecuteTemplate(w, "post.html", output); err != nil {
		h.InternalServerError500(w, r, err)
		return
	}
}
