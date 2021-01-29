package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salf       = "jdhf8fd7cd9dcudcd993xm45"
	singingKey = "8wh9adon2458cnldhr9393cm"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

// AuthService ...
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService ...
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser ...
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GenerateToken gen JWT-token
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(singingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salf)))
}
