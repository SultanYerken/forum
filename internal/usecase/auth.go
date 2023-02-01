package usecase

import (
	"forum/internal/entity"
	"forum/internal/repository"
	"forum/utils"
	"log"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	repo repository.Authorization
}

func NewAuthUseCase(repo repository.Authorization) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (u *AuthUseCase) CreateUser(user entity.User) (int, error) {
	pass, err := u.generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = pass

	return u.repo.CreateUser(user)
}

func (u *AuthUseCase) SignIn(username, password string) (string, error) {
	var token uuid.UUID

	userFromDB, err := u.repo.GetUser(username)
	if err != nil {
		return "", utils.ErrWrongLogin
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(password))
	if err != nil {
		return "", utils.ErrWrongPass
	}

	token, err = uuid.NewV4()
	if err != nil {
		log.Println("failed to generate UUID: %v", err)
		return "", err
	}

	_, tokeninStr, err := u.repo.CreateSession(userFromDB.Id, token.String())
	if err != nil {
		return "", err
	}

	return tokeninStr, nil
}

// for Middleware...
func (u *AuthUseCase) ParseToken(token string) (entity.Session, error) {
	return u.repo.GetSession(token)
}

func (u *AuthUseCase) LogOutUser(token string) error {
	return u.repo.DeleteSession(token)
}

func (u *AuthUseCase) generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
