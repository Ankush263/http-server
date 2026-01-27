package app

import (
	"log"

	"github.com/Ankush263/http-server/internals/config"
	"github.com/Ankush263/http-server/internals/router"
	"github.com/Ankush263/http-server/internals/server"
)

func Run() {
	cfg := config.Load()

	r := router.New()

	srv := server.New(cfg.Port, r)

	log.Printf("Server running on port %s\n: ", cfg.Port)
	log.Fatal(srv.Start())
}
