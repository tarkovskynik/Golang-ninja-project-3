package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tarkovskynik/Golang-ninja-project-3/internal/config"
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/db/mongo"
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/repository"
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/server"
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/services"
	"github.com/tarkovskynik/Golang-ninja-project-3/internal/transport/rest"

	"github.com/joho/godotenv"
)

func main() {
	cfg, err := initConfig("configs", "config")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := mongo.NewMongoDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := services.NewService(repo)
	handlers := rest.NewHandler(service)

	srv := server.NewServer(&http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:        handlers.Init(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	})

	go func() {
		log.Fatal(srv.Run())
	}()

	log.Println("Starting server on port 8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("Server shutting down...")

	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Failed to shutdown server: ", err)
	}

	if err = db.Disconnect(ctx); err != nil {
		log.Fatal("Failed to close database: ", err)
	}

	fmt.Println("Server stopped")
}

func initConfig(folder, file string) (*config.Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	cfg, err := config.New(folder, file)
	if err != nil {
		return nil, err
	}
	cfg.MongoDB.Password = os.Getenv("MONGODB_PASSWORD")

	return cfg, nil
}
