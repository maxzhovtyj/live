package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type MeetingService struct {
	rooms   map[string]*Room
	roomsMX sync.RWMutex
}

func NewMeetingService() Meeting {
	return &MeetingService{
		rooms: make(map[string]*Room),
	}
}

func (m *MeetingService) GetRoom(id string) (*Room, error) {
	m.roomsMX.RLock()
	r, ok := m.rooms[id]
	m.roomsMX.RUnlock()

	if ok && r != nil {
		return r, nil
	}

	return nil, fmt.Errorf("no such room")
}

func (m *MeetingService) NewRoom(id string) {
	m.roomsMX.Lock()
	defer m.roomsMX.Unlock()

	m.rooms[id] = &Room{
		ID:           id,
		participants: make(map[int32]*MeetParticipant),
	}
}

type Room struct {
	ID string

	participants   map[int32]*MeetParticipant
	participantsMX sync.RWMutex
}

func (r *Room) Subscribe(uid int32, conn *websocket.Conn) *MeetParticipant {
	r.participantsMX.Lock()
	defer r.participantsMX.Unlock()

	log.Println("connecting user to room", uid, r.ID)

	r.participants[uid] = &MeetParticipant{
		ID:       uid,
		Messages: make(chan BroadcastMessage),

		conn: conn,
	}

	return r.participants[uid]
}

func (r *Room) Unsubscribe(uid int32) {
	r.participantsMX.Lock()
	defer r.participantsMX.Unlock()

	close(r.participants[uid].Messages)

	delete(r.participants, uid)
}

func (r *Room) Publish(uid int32, msg BroadcastMessage) {
	r.participantsMX.RLock()
	defer r.participantsMX.RUnlock()

	for id, p := range r.participants {
		if id == uid {
			continue
		}

		p.Messages <- msg
	}
}

type BroadcastMessage struct {
	Message map[string]interface{}
	RoomID  string
}

type MeetParticipant struct {
	ID       int32
	Messages chan BroadcastMessage

	conn *websocket.Conn
}

func (mp *MeetParticipant) Write(v any) error {
	return mp.conn.WriteJSON(v)
}

func (mp *MeetParticipant) ReadJSON(v any) error {
	return mp.conn.ReadJSON(v)
}
