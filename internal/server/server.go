package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/Fam-Go/internal/middleware"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	router := gin.Default()
	router.Use(middleware.LoggingMiddleware)
	InitRoutes(router)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	return &Server{httpServer: httpServer}
}

func (s *Server) Start() error {
	go func() {
		log.Println("Starting server on port 8080...")
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port 8080: %v\n", err)
		}
	}()
	return waitForShutdown(s.httpServer)
}

func waitForShutdown(server *http.Server) error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}
