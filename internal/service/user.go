package service

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/maxzhovtyj/live/internal/models"
	"github.com/maxzhovtyj/live/internal/storage"
	"time"
)

const (
	signingKey = "123sdakm34l@@!sss"
	salt       = "12czc,xlc,ll122##@1"
)

type UserService interface {
	CreateUser(user models.User) error
	GenerateTokens(email string, password string) (string, error)
	ParseToken(token string) (int32, error)
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int32
}

type service struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) UserService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateUser(user models.User) error {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.User.Create(user)
}

func (s *service) GenerateTokens(email, password string) (string, error) {
	selectedUser, err := s.repo.User.GetAuthorizedUser(email, s.generatePasswordHash(password))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("користувача не знайдено")
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		selectedUser.ID,
	})

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *service) ParseToken(accessToken string) (int32, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (s *service) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
