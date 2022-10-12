package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	todo "http-repository-apiv1"
	"http-repository-apiv1/internal/handler"
	repository2 "http-repository-apiv1/internal/repository"
	"http-repository-apiv1/internal/service"
	"http-repository-apiv1/pkg/zaplogger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error init env config %s", err.Error())
	}

	// init logger
	logger, err := zaplogger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	db, err := repository2.NewPostgresDB(
		repository2.Config{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			DBName:   os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
			SSLMode:  os.Getenv("DB_SSL_MODE"),
		},
		logger,
	)
	if err != nil {
		log.Fatalf("error init db %s", err.Error())
	}

	repos := repository2.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err = srv.Run(os.Getenv("PORT"), handlers.InitRoutes(), logger); err != nil {
		log.Fatalf("error start server %s", err.Error())
	}
}
