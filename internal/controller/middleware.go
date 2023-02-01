package controller

import (
	"context"
	"forum/internal/entity"
	"log"
	"net/http"
)

const ctxKeyUser = "session_token"

func (h *Handler) authentificate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, user, err := h.getSessionFromCookie(w, r)
		if err != nil {
			h.Unauthorized401(w, r, err)
			return
		}
		log.Printf("Auth middleware %s OK", user.Username)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, session))) // ???
	})
}

func (h *Handler) getSessionFromCookie(w http.ResponseWriter, r *http.Request) (entity.Session, entity.User, error) {
	var session entity.Session
	var user entity.User

	cook, err := r.Cookie("session_token")
	if err != nil {
		return session, user, err
	}
	token := cook.Value

	session, err = h.usecases.ParseToken(token)
	if err != nil {
		return session, user, err
	}

	userId := session.UserId

	user, err = h.usecases.GetUser(userId)
	if err != nil {
		return session, user, err
	}
	return session, user, nil
}
