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
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "password0701",
		DbName:   "beletme",
		SSlmode:  "disable",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}

	repost := repository.NewRepository(db)
	newService := service.NewService(repost)
	newHandler := handler.NewHandler(newService)

	if err := server1.Run("8080", newHandler.InitRoutes()); err != nil {
		return
	}

}
