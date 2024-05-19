package service

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/maxzhovtyj/live/internal/models"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/storage"
	"time"
)

const (
	signingKey = "123sdakm34l@@!sss"
	salt       = "12czc,xlc,ll122##@1"
)

type tokenClaims struct {
	jwt.StandardClaims
	User db.User
}

type UserService struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() ([]db.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user models.User) error {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.User.Create(user)
}

func (s *UserService) GenerateTokens(email, password string) (string, error) {
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
		selectedUser,
	})

	accessToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *UserService) ParseToken(accessToken string) (db.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return db.User{}, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return db.User{}, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.User, nil
}

func (s *UserService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
