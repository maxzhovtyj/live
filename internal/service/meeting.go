package service

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type MeetingService struct {
	rooms   map[string]*Room
	roomsMX sync.RWMutex
}

func (m *MeetingService) GetRoom(id string) *Room {
	m.roomsMX.RLock()
	defer m.roomsMX.RUnlock()

	return m.rooms[id]
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
		Client:   conn,
		Messages: make(chan BroadcastMessage),
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

	for _, p := range r.participants {
		//if id == uid {
		//	continue
		//}

		log.Println("publish message to", p.ID)

		p.Messages <- msg
	}
}

type MeetParticipant struct {
	ID int32

	Client *websocket.Conn
	connMX sync.RWMutex

	Messages chan BroadcastMessage
}

func (mp *MeetParticipant) Write(v any) error {
	mp.connMX.Lock()
	defer mp.connMX.Unlock()

	return mp.Client.WriteJSON(v)
}

func (mp *MeetParticipant) ReadJSON(v any) error {
	mp.connMX.RLock()
	defer mp.connMX.RUnlock()

	return mp.Client.ReadJSON(v)
}

type BroadcastMessage struct {
	Message map[string]interface{}
	RoomID  string
}

func NewMeetingService() Meeting {
	return &MeetingService{
		rooms: make(map[string]*Room),
	}
}
