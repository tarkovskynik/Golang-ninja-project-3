package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tarkovskynik/Golang-ninja-project-3/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	cfg, err := initConfig("configs", "config")
	if err != nil {
		log.Fatalln(err)
	}

	addr := fmt.Sprintf("http://%s:%d", cfg.Server.Host, cfg.Server.Port)
	_, err = http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
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
