package app

import (
	"log"

	"github.com/Ankush263/http-server/internals/config"
	"github.com/Ankush263/http-server/internals/handler"
	"github.com/Ankush263/http-server/internals/middleware"
	"github.com/Ankush263/http-server/internals/router"
	"github.com/Ankush263/http-server/internals/server"
)

func Run() {
	cfg := config.Load()

	r := router.New()
	r.Handle("GET", "/health", handler.Health)
	r.Handle("POST", "/echo", handler.Echo)
	
	loggedRouter := middleware.Logger(r)

	srv := server.New(cfg.Port, loggedRouter)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
