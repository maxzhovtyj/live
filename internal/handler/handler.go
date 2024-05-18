package handler

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxzhovtyj/live/internal/service"
	"golang.org/x/time/rate"
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

	authorized := e.Group("/", h.Authorized)

	authorized.GET("", h.Index)
	authorized.GET("chat", h.Chat)

	authorized.GET("ws/chat", h.JoinChat)

	e.GET("/sign-in", h.SignInPage)
	e.POST("/sign-in", h.SignIn)
	e.POST("/sign-out", h.SignOut)

	e.GET("/sign-up", h.SignUpPage)
	e.POST("/sign-up", h.SignUp)

	return e
}
