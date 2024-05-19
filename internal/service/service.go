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
	GetAll() ([]db.User, error)
	CreateUser(user models.User) error
	GenerateTokens(email string, password string) (string, error)
	ParseToken(accessToken string) (db.User, error)
}

type Chat interface {
	NewChat(name string, ids ...int32) error
	NewMessage(cid, uid int32, msg string) error
	Join(cid int, cn *websocket.Conn, user db.User) (*Connection, *ChatRoom)
	GetRoom(cid int32) *ChatRoom
	GetRoomMessages(cid int) ([]db.GetConversationMessagesRow, error)
	GetUserConversations(id int32) ([]db.GetUserConversationsRow, error)
}

func New(repo *storage.Storage) *Service {
	return &Service{
		NewUserService(repo),
		NewChatService(repo),
	}
}
