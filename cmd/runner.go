package main

import (
	"log"
	"user-service/internal/controller"
	"user-service/internal/server/HTTPServer"
	"user-service/internal/server/config"
)

func main() {
	log.Println("Start programming application ...")

	userControl := controller.NewUserController()

	userControl.Router()

	conf := config.ConfigServer{
		Port: 8585,
		Host: "127.0.0.1",
	}

	httpServer := HTTPServer.NewWebServer(userControl.Mux, conf)

	httpServer.Start()
}
