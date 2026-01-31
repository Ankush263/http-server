package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(port string, handler http.Handler) *Server{
	return &Server{
		httpServer: &http.Server{
			Addr: ":" + port,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	go func() {
		log.Println("Server Started")
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<- quit

	log.Printf("Shutting down the server")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	return s.httpServer.Shutdown(ctx)
}
