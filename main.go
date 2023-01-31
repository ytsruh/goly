package main

import (
	"goly/model"
	"goly/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	model.Setup()
	server.SetupAndListen()
}
