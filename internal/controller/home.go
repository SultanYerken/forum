package controller

import (
	"forum/internal/entity"
	"forum/utils"
	"net/http"
)

type getPostsResponse struct {
	Posts []entity.Post
	User  entity.User
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}
	if r.URL.Path != "/" {
		h.NotFound404(w, r, utils.ErrPageNotFound404)
		return
	}

	// ts, err := template.ParseFiles(pageIndex)
	// if err != nil {
	// 	h.InternalServerError500(w, r, err)
	// 	return
	// }

	var output getPostsResponse
	var posts []entity.Post

	_, user, err := h.getSessionFromCookie(w, r)
	if err != nil {
		if err == http.ErrNoCookie {
			// fmt.Println("Home Page Unauthorized User")
		}
	}
	output.User = user

	r.ParseForm()
	category := r.Form["category"]

	if category == nil {
		posts, err := h.usecases.GetAllPost()
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
		output.Posts = posts
	} else {
		var categoriesId []int
		for i := range category {
			categoryId, err := h.usecases.GetCategoryId(category[i])
			if err != nil {
				h.BadRequest400(w, r, utils.ErrWrongCategory)
				return
			}
			categoriesId = append(categoriesId, categoryId)
		}

		postsId, err := h.usecases.GetPostsIdByCategory(categoriesId)
		if err != nil {
			h.InternalServerError500(w, r, err)
			return
		}

		for i := range postsId {
			post, err := h.usecases.GetPostById(postsId[i])
			if err != nil {
				h.InternalServerError500(w, r, err)
				return
			}
			posts = append(posts, post)
		}
		output.Posts = posts
	}

	// err = ts.Execute(w, output)
	// if err != nil {
	// 	h.InternalServerError500(w, r, err)
	// 	return
	// }

	if err := h.tmpl.ExecuteTemplate(w, "index.html", output); err != nil {
		h.InternalServerError500(w, r, err)
		return
	}
}
