package service

import "github.com/maxzhovtyj/live/internal/storage"

type Service struct {
	UserService
}

func New(repo *storage.Storage) *Service {
	return &Service{NewUserService(repo)}
}
