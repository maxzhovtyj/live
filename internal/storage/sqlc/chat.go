package sqlc

import (
	"context"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"time"
)

type ChatStorage struct {
	q *db.Queries
}

func (c *ChatStorage) AddUsersToConversation(cid int32, ids ...int32) error {
	for _, i := range ids {
		err := c.q.AddConversationParticipant(context.Background(), db.AddConversationParticipantParams{
			ConversationID: cid,
			UserID:         i,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ChatStorage) CreateConversation(name string) (int32, error) {
	return c.q.InsertConversation(context.Background(), name)
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
