package main

import (
	"log"

	web "github.com/Yosh11/exemple_gin"
	"github.com/Yosh11/exemple_gin/pkg/handler"
	"github.com/Yosh11/exemple_gin/pkg/repository"
	"github.com/Yosh11/exemple_gin/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(web.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while starting server: %s ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
