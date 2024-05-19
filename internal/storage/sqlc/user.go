package sqlc

import (
	"context"
	"github.com/maxzhovtyj/live/internal/models"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"time"
)

type UserStorage struct {
	q *db.Queries
}

func (u *UserStorage) GetAll() ([]db.User, error) {
	return u.q.GetAll(context.Background())
}

func (u *UserStorage) GetAuthorizedUser(email, passwordHash string) (db.User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	user, err := u.q.GetAuthorizedUser(ctx, db.GetAuthorizedUserParams{
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserStorage) Get(id int32) (db.User, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	user, err := u.q.GetUser(ctx, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewUserStorage(conn db.DBTX) *UserStorage {
	return &UserStorage{
		q: db.New(conn),
	}
}

func (u *UserStorage) Create(user models.User) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	_, err := u.q.CreateUser(ctx, db.CreateUserParams{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PasswordHash: user.Password,
	})
	if err != nil {
		return err
	}

	return nil
}
