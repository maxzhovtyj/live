package service

import (
	"github.com/gorilla/websocket"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/storage"
	"sync"
	"time"
)

type ChatService struct {
	chatRooms   map[int32]*ChatRoom
	chatRoomsMX sync.RWMutex

	repo *storage.Storage
}

func (c *ChatService) NewMessage(cid, uid int32, msg string) error {
	return c.repo.Chat.InsertMessageIntoConversation(cid, uid, msg)
}

type Message struct {
	FirstName string
	LastName  string
	String    string
	Time      time.Time
}

type ChatRoom struct {
	ConversationID int32

	connections   map[int32]*Connection
	connectionsMX sync.RWMutex

	Upgrader *websocket.Upgrader
}

func (cr *ChatRoom) Publish(m Message) error {
	cr.connectionsMX.RLock()
	cr.connectionsMX.RUnlock()

	for _, conn := range cr.connections {
		conn.Messages <- m
	}

	return nil
}

func NewChatService(repo *storage.Storage) *ChatService {
	return &ChatService{
		chatRooms: make(map[int32]*ChatRoom),
		repo:      repo,
	}
}

type Connection struct {
	User     db.User
	Messages chan Message
	Conn     *websocket.Conn
}

func (c *ChatService) Join(cid int, cn *websocket.Conn, user db.User) (*Connection, *ChatRoom) {
	room := c.GetRoom(int32(cid))

	return room.Join(user.ID, cn, user), room
}

func (c *ChatService) GetRoom(cid int32) *ChatRoom {
	c.chatRoomsMX.RLock()
	cr, ok := c.chatRooms[cid]
	c.chatRoomsMX.RUnlock()

	if ok && cr != nil {
		return cr
	}

	c.chatRoomsMX.Lock()
	defer c.chatRoomsMX.Unlock()

	cr, ok = c.chatRooms[cid]
	if ok && cr != nil {
		return cr
	}

	cr = &ChatRoom{
		ConversationID: cid,
		connections:    map[int32]*Connection{},
	}
	c.chatRooms[cid] = cr

	return cr
}

func (cr *ChatRoom) Join(id int32, c *websocket.Conn, user db.User) *Connection {
	cr.connectionsMX.Lock()
	cn := &Connection{
		User:     user,
		Messages: make(chan Message),
		Conn:     c,
	}
	cr.connections[id] = cn
	cr.connectionsMX.Unlock()

	return cn
}

func (cr *ChatRoom) Leave(id int32) {
	cr.connectionsMX.Lock()
	defer cr.connectionsMX.Unlock()

	c := cr.connections[id]
	close(c.Messages)

	delete(cr.connections, id)
}

func (c *ChatService) GetRoomMessages(cid int) ([]db.GetConversationMessagesRow, error) {
	return c.repo.GetConversationMessages(int32(cid))
}
