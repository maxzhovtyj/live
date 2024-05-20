package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/live/internal/pkg/templates/video"
	"github.com/maxzhovtyj/live/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms = service.NewVideoRoom()

type response struct {
	RoomID string `json:"room_id"`
}

type BroadcastMessage struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan BroadcastMessage)

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range AllRooms.Participants[msg.RoomID] {
			if client.Conn != msg.Client {
				client.Mutex.Lock()

				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					log.Println(err)
					client.Conn.Close()
				}

				client.Mutex.Unlock()
			}
		}
	}
}

func (h *Handler) VideoRoom(ctx echo.Context) error {
	return video.VideoRoom().Render(context.Background(), ctx.Response().Writer)
}

var videoUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) CreateRoomRequestHandler(ctx echo.Context) error {
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")

	roomID := AllRooms.CreateRoom()

	return json.NewEncoder(ctx.Response()).Encode(response{RoomID: roomID})
}

func (h *Handler) JoinRoomRequestHandler(ctx echo.Context) error {
	roomID := ctx.QueryParam("roomID")

	if roomID == "" {
		return fmt.Errorf("roomID is missing, unable to join the call")
	}

	ws, err := videoUpgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
	if err != nil {
		log.Println("Unable to upgrade http to websocket", err)
		return err
	}

	AllRooms.InsertIntoRoom(roomID, false, ws)

	go broadcaster()

	for {
		var msg BroadcastMessage

		err = ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println(err)
			return err
		}

		msg.Client = ws
		msg.RoomID = roomID

		broadcast <- msg
	}
}
