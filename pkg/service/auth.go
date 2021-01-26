package service

import (
	"crypto/sha1"
	"fmt"

	web "github.com/Yosh11/exemple_gin"
	"github.com/Yosh11/exemple_gin/pkg/repository"
)

const (
	salf = "jdhf8fd7cd9dcudcd993xm45"
)

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ...
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ...
func (s *AuthService) CreateUser(user web.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salf)))
}
