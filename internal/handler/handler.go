package handler

import (
	"context"
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"
	"github.com/maxzhovtyj/live/internal/pkg/templates/components"
	"github.com/maxzhovtyj/live/internal/service"
	"golang.org/x/time/rate"
	"strconv"
)

var publicFolder = flag.String("public", "/Users/maksymzhovtaniuk/Desktop/univer4.2/диплом/live/internal/public/*.html", "Public folder path")

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Init() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	e.Static("/static/", "/Users/maksymzhovtaniuk/Desktop/univer4.2/диплом/live/static")

	authorized := e.Group("/", h.Authorized)

	authorized.GET("", h.Index)

	authorized.GET("chat", h.Chat)
	authorized.GET("ws/chat", h.JoinChat)
	authorized.GET("conversations", h.Conversations)

	authorized.GET("modal", h.newChatModal)
	authorized.POST("new-chat", h.newChat)

	authorized.GET("video", h.VideoRoom)
	authorized.GET("create-room", h.CreateRoomRequestHandler)
	authorized.GET("join-room", h.JoinRoomRequestHandler)

	e.GET("/sign-in", h.SignInPage)
	e.POST("/sign-in", h.SignIn)
	e.POST("/sign-out", h.SignOut)

	e.GET("/sign-up", h.SignUpPage)
	e.POST("/sign-up", h.SignUp)

	return e
}

func (h *Handler) newChatModal(ctx echo.Context) error {
	users, err := h.s.User.GetAll()
	if err != nil {
		return err
	}

	curr := h.getUserFromContext(ctx)
	var other []db.User

	for _, u := range users {
		if u.ID != curr.ID {
			other = append(other, u)
		}
	}

	return components.NewConversation(other).Render(context.Background(), ctx.Response().Writer)
}

func (h *Handler) newChat(ctx echo.Context) error {
	name := ctx.FormValue("name")
	user, err := strconv.Atoi(ctx.FormValue("user"))
	if err != nil {
		return err
	}

	u := h.getUserFromContext(ctx)

	err = h.s.Chat.NewChat(name, []int32{int32(user), u.ID}...)
	if err != nil {
		return err
	}

	ctx.Response().Header().Set("HX-Redirect", "/chat")

	return nil

}
