package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // ...
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/handler"
	"github.com/Yosh11/exemple_gin/pkg/repository"
	"github.com/Yosh11/exemple_gin/pkg/service"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewpostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Errorf("error occurred while starting server: %s", err.Error())
		}
	}()

	log.Infoln("TodoApp Started")

	// Graceful shutdown server and db with timeout
	graceful(srv, db, 5*time.Second)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func graceful(srv *todo.Server, db *sqlx.DB, timeout time.Duration) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := db.Close(); err != nil {
		log.Errorf("error occurred on db connection close: %s", err.Error())
	} else {
		log.Infoln("db closed")
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("error occurred on server shutting down: %s", err.Error())
	} else {
		log.Infof("TodoApp ShutDown with timeout: %v", timeout)
	}
}
