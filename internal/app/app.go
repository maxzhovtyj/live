package app

import (
	"github.com/maxzhovtyj/live/internal/handler"
	"github.com/maxzhovtyj/live/internal/service"
	"github.com/maxzhovtyj/live/internal/storage"
	"github.com/maxzhovtyj/live/pkg/db/postgres"
	"log"
)

func Run() {
	dbConn, err := postgres.NewConn()
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
