package utils

import "errors"

var (
	ErrBadRequest400         = errors.New("Bad Request")
	ErrUnauthorized401       = errors.New("Unauthorized")
	ErrPageNotFound404       = errors.New("Page Not Found")
	ErrMethodNodAllowed405   = errors.New("Method Not Allowed")
	ErrInternalSeverError500 = errors.New("Internal Server Error")

	ErrEmptyFields = errors.New("Nothing entered in the fields")

	ErrPass1Pass2 = errors.New("Password 1 and password 2 do not match")

	ErrorNameExist    = errors.New("Username exist, choose another username")
	ErrorEmailExist   = errors.New("Email is used by another user, choose another email")
	ErrorSessionExist = errors.New("Session with this user exist")

	ErrWrongLogin = errors.New("Wrong Login or you are not registered")
	ErrWrongPass  = errors.New("Wrong Password")

	ErrWrongCategory = errors.New("Сategory entered incorrectly")

	ErrPostNotExist = errors.New("Post Not Exist")

	ErrCommentNotExist = errors.New("Comment Not Exist")
)
