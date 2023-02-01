package controller

import (
	"forum/internal/entity"
	"forum/utils"
	"log"
	"net/http"
	"time"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ts, err := template.ParseFiles(pageSignUp)
		// if err != nil {
		// 	h.InternalServerError500(w, r, err)
		// 	return
		// }
		// err = ts.Execute(w, nil)
		// if err != nil {
		// 	h.InternalServerError500(w, r, err)
		// 	return
		// }

		if err := h.tmpl.ExecuteTemplate(w, "signup.html", nil); err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
	} else if r.Method == http.MethodPost {
		var input entity.User

		input.Username = r.FormValue("username")
		input.Email = r.FormValue("email")
		input.Password = r.FormValue("password")
		password2 := r.FormValue("password2")

		arrayFields := []string{input.Username, input.Email, input.Password, password2}
		if chekEmptyField(arrayFields) {
			h.BadRequest400(w, r, utils.ErrEmptyFields)
			return
		}

		if input.Password != password2 {
			h.BadRequest400(w, r, utils.ErrPass1Pass2)
			return
		}

		id, err := h.usecases.CreateUser(input)
		if err != nil {
			h.ChekErrors(w, r, err)
			return
		}

		log.Printf("user %s registered, id = %d \n", input.Username, id)

		http.Redirect(w, r, "/signin", http.StatusSeeOther)

	} else {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}
}

type signInInput struct {
	Username string
	Password string
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// ts, err := template.ParseFiles(pageSignIn)
		// if err != nil {
		// 	h.InternalServerError500(w, r, err)
		// 	return
		// }
		// err = ts.Execute(w, nil)
		// if err != nil {
		// 	h.InternalServerError500(w, r, err)
		// 	return
		// }

		if err := h.tmpl.ExecuteTemplate(w, "signin.html", nil); err != nil {
			h.InternalServerError500(w, r, err)
			return
		}
	} else if r.Method == http.MethodPost {

		var input signInInput

		input.Username = r.FormValue("username")
		input.Password = r.FormValue("password")

		arrayFields := []string{input.Username, input.Password}
		if chekEmptyField(arrayFields) {
			h.BadRequest400(w, r, utils.ErrEmptyFields)
			return
		}

		token, err := h.usecases.SignIn(input.Username, input.Password)
		if err != nil {
			h.ChekErrors(w, r, err)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token", //"session_token",
			Value:   token,
			Expires: time.Now().Add(12 * time.Hour),
		})

		log.Printf("user %s authorized \n", input.Username)

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}
}

func (h *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.MethodNotAllowed405(w, r, utils.ErrMethodNodAllowed405)
		return
	}

	session, user, _ := h.getSessionFromCookie(w, r)
	token := session.Token

	err := h.usecases.LogOutUser(token)
	if err != nil {
		h.InternalServerError500(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		// Value:   "",
		Expires: time.Now(),
	})

	log.Printf("user %s logged out \n", user.Username)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func chekEmptyField(array []string) bool {
	var count int
	for i := range array {
		for _, w := range array[i] {
			if w == ' ' {
				count++
			}
		}
		if count == len(array[i]) {
			return true
		}
		count = 0
	}
	return false
}
