package main

import (
	"log"

	web "github.com/Yosh11/exemple_gin"
	"github.com/Yosh11/exemple_gin/pkg/handler"
	"github.com/Yosh11/exemple_gin/pkg/repository"
	"github.com/Yosh11/exemple_gin/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(web.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while starting server: %s ", err.Error())
	}
}
