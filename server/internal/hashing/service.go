package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Hash(password string) (string, error)
	Verify(password string, hash string) bool
}

type service struct {
}

func (s service) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (s service) Verify(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func NewService() Service {
	return &service{}
}
