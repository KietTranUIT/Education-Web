package main

import (
	"log"
	"user-service/internal/controller"
	"user-service/internal/core/service"
	"user-service/internal/infra/configDB"
	"user-service/internal/infra/repository"
	"user-service/internal/server/HTTPServer"
	"user-service/internal/server/config"
)

func main() {
	log.Println("Start programming application ...")

	confDB := configDB.ConfigDatabase{
		Driver:   "mysql",
		Username: "kiettran",
		Password: "KietTran@2003",
		Host:     "127.0.0.1",
		Port:     3306,
		Dbname:   "EDUCATION",
	}

	db := repository.NewDB(confDB)

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userControl := controller.NewUserController(userService)

	userControl.Router()

	conf := config.ConfigServer{
		Port: 8585,
		Host: "127.0.0.1",
	}

	httpServer := HTTPServer.NewWebServer(userControl.Mux, conf)

	httpServer.Start()
}
