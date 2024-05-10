package ui

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxzhovtyj/live/internal/pkg/template"
	"golang.org/x/time/rate"
	"net/http"
)

var publicFolder = flag.String("public", "/Users/maksymzhovtaniuk/Desktop/univer4.2/диплом/live/internal/public/*.html", "Public folder path")

func Run() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	template.NewTemplateRenderer(e, *publicFolder)
	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index", nil)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
