package controller

import (
	"forum/internal/usecase"
	"html/template"
	"net/http"
)

const (
	pageBadRequest400          = "./templates/html/400.html"
	pageUnauthorized401        = "./templates/html/401.html"
	pageNotFound404            = "./templates/html/404.html"
	pageMethodNotAllowed405    = "./templates/html/405.html"
	pageInternalServerError500 = "./templates/html/500.html"

	pageIndex      = "./templates/html/index.html"
	pagePost       = "./templates/html/post.html"
	pageCreatePost = "./templates/html/create_post.html"
	pageUser       = "./templates/html/user.html"
	pageSignUp     = "./templates/html/signup.html"
	pageSignIn     = "./templates/html/signin.html"
)

type Handler struct {
	usecases *usecase.UseCase
	tmpl     *template.Template
}

func NewHandler(usecases *usecase.UseCase) *Handler {
	return &Handler{
		usecases: usecases,
		tmpl:     template.Must(template.ParseGlob("./templates/html/*.html")),
	}
}

func (h *Handler) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Home)

	mux.HandleFunc("/signup", h.SignUp) // POST
	mux.HandleFunc("/signin", h.SignIn) // POST
	mux.Handle("/logout", h.authentificate(http.HandlerFunc(h.LogOut)))
	mux.Handle("/user", h.authentificate(http.HandlerFunc(h.GetUserProfile)))

	mux.Handle("/posts/", h.authentificate(http.HandlerFunc(h.CreatePost))) // POST
	mux.HandleFunc("/posts", h.GetPostById)                                 // GET

	mux.Handle("/comment/", h.authentificate(http.HandlerFunc(h.CreateComment))) // POST

	mux.Handle("/likes/", h.authentificate(http.HandlerFunc(h.Like)))       // POST
	mux.Handle("/dislikes/", h.authentificate(http.HandlerFunc(h.Dislike))) // POST

	fileServer := http.FileServer(http.Dir("./templates/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
