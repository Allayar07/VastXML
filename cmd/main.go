package main

import (
	"golang_vast/internal/handler"
	"golang_vast/internal/repository"
	"golang_vast/internal/server"
	"golang_vast/internal/service"
	"log"
)

func main() {
	server1 := new(server.Server)

	db, err := repository.NewPostgres(repository.Config{
		Host:     "192.168.1.61",
		Port:     "5454",
		Username: "tmtube_admin",
		Password: "22tmtubeadmin22",
		DbName:   "tmtube_db",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}

	repost := repository.NewRepository(db)
	newService := service.NewService(repost)
	newHandler := handler.NewHandler(newService)

	if err := server1.Run("8081", newHandler.InitRoutes()); err != nil {
		return
	}

}
