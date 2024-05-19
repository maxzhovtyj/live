package service

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Participant struct {
	Host  bool
	ID    string
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

type VideoRoom struct {
	Participants map[string][]*Participant
	mx           sync.RWMutex
}

func NewVideoRoom() *VideoRoom {
	return &VideoRoom{
		Participants: make(map[string][]*Participant),
	}
}

// Get all participant in a room
func (r *VideoRoom) Get(roomID string) []*Participant {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return r.Participants[roomID]
}

// CreateRoom creates a room
func (r *VideoRoom) CreateRoom() string {
	r.mx.Lock()
	defer r.mx.Unlock()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.Participants[roomID] = []*Participant{}

	return roomID
}

// InsertIntoRoom join a room handler
func (r *VideoRoom) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.mx.Lock()
	defer r.mx.Unlock()

	clientID := uuid.New().String()
	incomingParticipant := &Participant{host, clientID, conn, sync.Mutex{}}

	r.Participants[roomID] = append(r.Participants[roomID], incomingParticipant)
}

// DeleteRoom delete a room
func (r *VideoRoom) DeleteRoom(roomID string) {
	r.mx.Lock()
	defer r.mx.Unlock()

	delete(r.Participants, roomID)
}
