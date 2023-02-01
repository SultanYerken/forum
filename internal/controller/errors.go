package controller

import (
	"errors"
	"forum/utils"
	"html/template"
	"log"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Println("error handler:", message, "statuscode:", statusCode)
	http.Error(w, message, statusCode)
}

func (h *Handler) BadRequest400(w http.ResponseWriter, r *http.Request, message error) {
	ts, err := template.ParseFiles(pageBadRequest400)
	if err != nil {
		log.Println("parse template error:", err)
		h.InternalServerError500(w, r, err)
		return
	}
	output := errorResponse{
		Message: message.Error(),
	}
	log.Println("error:", message)
	w.WriteHeader(http.StatusBadRequest)
	ts.Execute(w, output)
}

func (h *Handler) Unauthorized401(w http.ResponseWriter, r *http.Request, message error) {
	ts, err := template.ParseFiles(pageUnauthorized401)
	if err != nil {
		log.Println("parse template error:", err)
		h.InternalServerError500(w, r, err)
		return
	}
	output := errorResponse{
		Message: message.Error(),
	}
	log.Println("error:", message)
	w.WriteHeader(http.StatusUnauthorized)
	ts.Execute(w, output)
}

func (h *Handler) NotFound404(w http.ResponseWriter, r *http.Request, message error) {
	ts, err := template.ParseFiles(pageNotFound404)
	if err != nil {
		log.Println("parse template error:", err)
		h.InternalServerError500(w, r, err)
		return
	}
	log.Println("error:", message)
	w.WriteHeader(http.StatusNotFound)
	ts.Execute(w, nil)
}

func (h *Handler) MethodNotAllowed405(w http.ResponseWriter, r *http.Request, message error) {
	ts, err := template.ParseFiles(pageMethodNotAllowed405)
	if err != nil {
		log.Println("parse template error:", err)
		h.InternalServerError500(w, r, err)
		return
	}
	log.Println("error:", message)
	w.WriteHeader(http.StatusMethodNotAllowed)
	ts.Execute(w, nil)
}

func (h *Handler) InternalServerError500(w http.ResponseWriter, r *http.Request, message error) {
	ts, err := template.ParseFiles(pageInternalServerError500)
	if err != nil {
		log.Println("parse template error:", err)
		h.InternalServerError500(w, r, err)
		return
	}
	output := errorResponse{
		Message: message.Error(),
	}
	log.Println("error:", message)
	w.WriteHeader(http.StatusInternalServerError)
	ts.Execute(w, output)
}

func (h *Handler) ChekErrors(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, utils.ErrBadRequest400) || errors.Is(err, utils.ErrEmptyFields) || errors.Is(err, utils.ErrWrongLogin) || errors.Is(err, utils.ErrWrongPass) || errors.Is(err, utils.ErrorNameExist) || errors.Is(err, utils.ErrorEmailExist) {
		h.BadRequest400(w, r, err)
		return
	} else if errors.Is(err, utils.ErrUnauthorized401) {
		h.Unauthorized401(w, r, err)
		return
	} else if errors.Is(err, utils.ErrPageNotFound404) {
		h.NotFound404(w, r, err)
		return
	} else if errors.Is(err, utils.ErrMethodNodAllowed405) {
		h.MethodNotAllowed405(w, r, err)
		return
	} else {
		h.InternalServerError500(w, r, err)
		return
	}
}
