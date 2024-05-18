package sqlc

import (
	"context"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"time"
)

type ChatStorage struct {
	q *db.Queries
}

func NewChatStorage(conn db.DBTX) *ChatStorage {
	return &ChatStorage{
		q: db.New(conn),
	}
}

func (c *ChatStorage) InsertMessageIntoConversation(cid, sender int32, msg string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	return c.q.InsertMessageIntoConversation(ctx, db.InsertMessageIntoConversationParams{
		ConversationID: cid,
		SenderID:       sender,
		Body:           msg,
	})
}

func (c *ChatStorage) GetConversation(id int32) (db.Conversation, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	return c.q.GetConversation(ctx, id)
}

func (c *ChatStorage) GetConversationMessages(id int32) ([]db.GetConversationMessagesRow, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	return c.q.GetConversationMessages(ctx, id)
}

func (c *ChatStorage) GetUserConversations(id int32) ([]db.GetUserConversationsRow, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	return c.q.GetUserConversations(ctx, id)
}
