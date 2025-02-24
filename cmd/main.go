package main

import (
	"log"

	"github.com/adityatresnobudi/restapi/config"
	"github.com/adityatresnobudi/restapi/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("err loading .env file: %s\n", err.Error())
	}
}

// @title Transaction API
// @version 1.0
// @description Transaction API written in Go
// @BasePath /
func main() {
	cfg := config.NewConfig()

	s := server.NewServer(cfg)

	s.Run()
}
