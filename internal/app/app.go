package app

import (
	"flag"
	"github.com/maxzhovtyj/live/internal/config"
	"github.com/maxzhovtyj/live/internal/handler"
	"github.com/maxzhovtyj/live/internal/service"
	"github.com/maxzhovtyj/live/internal/storage"
	"github.com/maxzhovtyj/live/pkg/db/postgres"
	"log"
)

func Run() {
	flag.Parse()

	dbConn, err := postgres.NewConn(config.Get().DBConnection)
	if err != nil {
		log.Fatal(err)
	}

	appStorage := storage.New(dbConn)
	appService := service.New(appStorage)
	appHandler := handler.New(appService)

	server := appHandler.Init()

	if err = server.Start(":6789"); err != nil {
		log.Fatal(err)
	}
}
