package storage

import (
	"github.com/maxzhovtyj/live/internal/models"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/storage/sqlc"
)

type Storage struct {
	User
	Chat
}

type User interface {
	Create(user models.User) error
	Get(id int32) (db.User, error)
	GetAuthorizedUser(email, passwordHash string) (db.User, error)
}

type Chat interface {
	InsertMessageIntoConversation(cid, sender int32, msg string) error
	GetConversation(id int32) (db.Conversation, error)
	GetConversationMessages(id int32) ([]db.GetConversationMessagesRow, error)
	GetUserConversations(id int32) ([]db.ConversationParticipant, error)
}

func New(conn db.DBTX) *Storage {
	return &Storage{
		User: sqlc.NewUserStorage(conn),
		Chat: sqlc.NewChatStorage(conn),
	}
}
