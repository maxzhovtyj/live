package service

import (
	"github.com/gorilla/websocket"
	"github.com/maxzhovtyj/live/internal/models"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/storage"
)

type Service struct {
	User
	Chat
}

type User interface {
	CreateUser(user models.User) error
	GenerateTokens(email string, password string) (string, error)
	ParseToken(accessToken string) (db.User, error)
}

type Chat interface {
	Join(cid int, cn *websocket.Conn, user db.User) (*Connection, *ChatRoom)
	GetRoom(cid int) *ChatRoom
	GetRoomMessages(cid int) ([]db.GetConversationMessagesRow, error)
}

func New(repo *storage.Storage) *Service {
	return &Service{
		NewUserService(repo),
		NewChatService(repo),
	}
}
