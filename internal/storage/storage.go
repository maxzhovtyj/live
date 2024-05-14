package storage

import (
	"github.com/maxzhovtyj/live/internal/models"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/storage/sqlc"
)

type Storage struct {
	User
}

type User interface {
	Create(user models.User) error
	Get(id int32) (db.User, error)
	GetAuthorizedUser(email, passwordHash string) (db.User, error)
}

func New(conn db.DBTX) *Storage {
	return &Storage{
		User: sqlc.NewUserStorage(conn),
	}
}
