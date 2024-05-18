package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/maxzhovtyj/live/internal/pkg/templates"
	"github.com/maxzhovtyj/live/internal/pkg/templates/components"
	"github.com/maxzhovtyj/live/internal/service"
	"log"
	"time"
)

func (h *Handler) Chat(ctx echo.Context) error {
	return templates.Chat().Render(context.Background(), ctx.Response().Writer)
}

var upgrader websocket.Upgrader

func (h *Handler) JoinChat(ctx echo.Context) error {
	cid := 1

	conn, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
	if err != nil {
		return err
	}

	chatConn, room := h.s.Chat.Join(cid, conn, h.getUserFromContext(ctx))
	defer room.Leave(chatConn.User.ID)

	messages, err := h.s.Chat.GetRoomMessages(cid)
	if err != nil {
		return err
	}

	messagesBuffer := bytes.NewBuffer(nil)

	for _, m := range messages {
		msg := components.Message(m.Concat.(string), m.CreatedAt.Time.Format(time.Kitchen), m.Body)
		err = msg.Render(context.Background(), messagesBuffer)
		if err != nil {
			return err
		}
	}

	err = conn.WriteMessage(websocket.TextMessage, messagesBuffer.Bytes())
	if err != nil {
		return err
	}

	closeCh := make(chan struct{})

	go h.wsReader(closeCh, chatConn, room)
	go h.wsWriter(chatConn)

	<-closeCh

	return nil
}

type ClientMessage struct {
	ChatMessage string `json:"chat-message"`
}

func (h *Handler) wsReader(cl chan struct{}, chatConn *service.Connection, room *service.ChatRoom) {
	defer close(cl)

	for {
		t, b, err := chatConn.Conn.ReadMessage()
		if t == websocket.CloseMessage || err != nil {
			return
		}

		if t == -1 {
			continue
		}

		var cm ClientMessage

		err = json.Unmarshal(b, &cm)
		if err != nil {
			log.Println(err)
			continue
		}

		err = room.Publish(service.Message{
			FirstName: chatConn.User.FirstName,
			LastName:  chatConn.User.LastName,
			String:    cm.ChatMessage,
			Time:      time.Now(),
		})
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func (h *Handler) wsWriter(chatConn *service.Connection) {
	for m := range chatConn.Messages {
		n := fmt.Sprintf("%s %s", m.FirstName, m.LastName)

		b := bytes.NewBuffer(nil)
		msg := components.Message(n, m.Time.Format(time.Kitchen), m.String)

		err := msg.Render(context.Background(), b)
		if err != nil {
			log.Println(err)
			return
		}

		err = chatConn.Conn.WriteMessage(websocket.TextMessage, b.Bytes())
		if err != nil {
			log.Println(err)
			return
		}
	}
}