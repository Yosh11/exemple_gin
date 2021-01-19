package main

import (
	"log"

	web "github.com/Yosh11/exemple_gin"
	"github.com/Yosh11/exemple_gin/pkg/handler"
)

func main() {
	srv := new(web.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while starting server: %s ", err.Error())
	}
}
