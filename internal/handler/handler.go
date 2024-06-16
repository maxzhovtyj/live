package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxzhovtyj/live/internal/config"
	"github.com/maxzhovtyj/live/internal/service"
	"golang.org/x/time/rate"
)

type Handler struct {
	s *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Init() *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	e.Static("/static/", config.Get().StaticDirPath)

	e.Use(middleware.Logger())

	authorized := e.Group("/", h.Authorized)

	authorized.GET("", h.Index)

	authorized.GET("chat", h.Chat)
	authorized.GET("ws/chat", h.JoinChat)
	authorized.GET("conversations", h.Conversations)

	authorized.GET("modal", h.newChatModal)
	authorized.POST("new-chat", h.newChat)

	authorized.POST("create-meeting", h.CreateRoomRequestHandler)

	authorized.GET("meeting", h.VideoRoom)
	authorized.GET("ws/join-room", h.JoinRoomRequestHandler)

	e.GET("/sign-in", h.SignInPage)
	e.POST("/sign-in", h.SignIn)
	e.POST("/sign-out", h.SignOut)

	e.GET("/sign-up", h.SignUpPage)
	e.POST("/sign-up", h.SignUp)

	return e
}
