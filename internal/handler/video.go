package handler

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/live/internal/pkg/templates/video"
	"github.com/maxzhovtyj/live/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (h *Handler) VideoRoom(ctx echo.Context) error {
	meetingID := ctx.QueryParam("id")

	c := getContext(ctx)

	if meetingID == "" {
		return video.MeetingPage(c).Render(context.Background(), ctx.Response().Writer)
	}

	return video.VideoRoom(c).Render(context.Background(), ctx.Response().Writer)
}

var videoUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) CreateRoomRequestHandler(ctx echo.Context) error {
	roomID := uuid.New()

	h.s.Meeting.NewRoom(roomID.String())

	ctx.Response().Header().Set("HX-Redirect", "/meeting?id="+roomID.String())

	return nil
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

	u := getContext(ctx).User

	room, err := h.s.Meeting.GetRoom(roomID)
	if err != nil {
		return err
	}

	sub := room.Subscribe(u.ID, ws)
	defer room.Unsubscribe(u.ID)

	go func() {
		for m := range sub.Messages {
			err = sub.Write(m.Message)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()

	for {
		var msg service.BroadcastMessage

		err = sub.ReadJSON(&msg.Message)
		if err != nil {
			log.Println(err)
			return err
		}

		room.Publish(u.ID, msg)
	}
}
