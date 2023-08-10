package HTTPServer

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/internal/server/config"
)

type HTTP_Server interface {
	Start()
	Stop()
}

type WebServer struct {
	port   uint
	host   string
	server http.Server
}

func NewWebServer(mux *http.ServeMux, conf config.ConfigServer) HTTP_Server {
	return &WebServer{
		port: conf.Port,
		host: conf.Host,
		server: http.Server{
			Addr:    conf.GetURL(),
			Handler: mux,
		},
	}
}

func (server *WebServer) Start() {
	go func() {
		if err := server.server.ListenAndServe(); err != nil {
			log.Println("Server error: ", err.Error())
		}
	}()
	server.Stop()
}

func (server *WebServer) Stop() {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	sig := <-signalCh
	log.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.server.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shut down server: ", err.Error())
	}
}
