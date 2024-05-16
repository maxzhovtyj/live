package video

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

type Room struct {
	participants map[string][]*Participant
	mx           sync.RWMutex
}

// Init initialize room map
func (r *Room) Init() {
	r.participants = make(map[string][]*Participant)
}

// Get all participant in a room
func (r *Room) Get(roomID string) []*Participant {
	r.mx.RLock()
	defer r.mx.RUnlock()

	return r.participants[roomID]
}

// CreateRoom creates a room
func (r *Room) CreateRoom() string {
	r.mx.Lock()
	defer r.mx.Unlock()

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)
	r.participants[roomID] = []*Participant{}

	return roomID
}

// InsertIntoRoom join a room handler
func (r *Room) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.mx.Lock()
	defer r.mx.Unlock()

	clientID := uuid.New().String()
	incomingParticipant := &Participant{host, clientID, conn, sync.Mutex{}}

	r.participants[roomID] = append(r.participants[roomID], incomingParticipant)
}

// DeleteRoom delete a room
func (r *Room) DeleteRoom(roomID string) {
	r.mx.Lock()
	defer r.mx.Unlock()

	delete(r.participants, roomID)
}
