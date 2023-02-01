package repository

import (
	"database/sql"
	"forum/internal/entity"
	"forum/utils"
	"log"
)

type AuthDataBase struct {
	db *sql.DB
}

func NewAuthDataBase(db *sql.DB) *AuthDataBase {
	return &AuthDataBase{db: db}
}

func (r *AuthDataBase) CreateUser(user entity.User) (int, error) {
	var id int

	if userExist(r.db, user.Username) {
		return 0, utils.ErrorNameExist
	}
	if emailExist(r.db, user.Email) {
		return 0, utils.ErrorEmailExist
	}

	query := `INSERT INTO user (username,email,password) VALUES($1,$2,$3) RETURNING id`
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthDataBase) GetUser(username string) (entity.User, error) {
	var user entity.User
	query := `SELECT id, userName, email, password FROM user WHERE userName = $1`
	row := r.db.QueryRow(query, username)
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		log.Println("repository GetUser", err)
		return user, err
	}
	return user, err // zhashkevish fa cosi
}

func (r *AuthDataBase) CreateSession(userId int, token string) (int, string, error) {
	session, err := sessionWithUserExist(r.db, userId)
	if err != nil {
		return 0, "", err
	}
	if session.Token != "" {
		err := updateSession(r.db, session.Token, token)
		if err != nil {
			return 0, "", err
		}
		return session.Id, token, nil
	}

	query := `INSERT INTO session (userId, token) VALUES($1,$2)`
	res, err := r.db.Exec(query, userId, token)
	if err != nil {
		return 0, "", err
	}
	sessionId, err := res.LastInsertId()
	if err != nil {
		return 0, "", err
	}
	return int(sessionId), token, nil
}

// for middleware...
func (r *AuthDataBase) GetSession(token string) (entity.Session, error) {
	var session entity.Session
	query := `SELECT id, userId, token FROM session WHERE token = $1`
	err := r.db.QueryRow(query, token).Scan(&session.Id, &session.UserId, &session.Token)
	if err != nil {
		return session, err
	}
	return session, nil
}

func (r *AuthDataBase) DeleteSession(token string) error {
	query := `DELETE FROM session WHERE token = $1`
	_, err := r.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}

func updateSession(db *sql.DB, sessionToken string, newToken string) error {
	query := `UPDATE session SET token = $1 WHERE token = $2`
	_, err := db.Exec(query, newToken, sessionToken)
	if err != nil {
		return err
	}
	return nil
}

func userExist(db *sql.DB, username string) bool {
	query := `SELECT userName FROM user WHERE userName = $1`
	err := db.QueryRow(query, username).Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
		}
		return false
	}
	return true
}

func emailExist(db *sql.DB, email string) bool {
	query := `SELECT email FROM user WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&email)
	if err != nil {
		if err != sql.ErrNoRows {
		}
		return false
	}
	return true
}

func sessionWithUserExist(db *sql.DB, userId int) (entity.Session, error) {
	var sess entity.Session
	query := `SELECT id, userId, token FROM session WHERE userId = $1`
	err := db.QueryRow(query, userId).Scan(&sess.Id, &sess.UserId, &sess.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return sess, nil // нужно nil чтобы он дальше создавал сессию,а не отправлял ошибку
		}
		return sess, err
	}
	return sess, nil
}
