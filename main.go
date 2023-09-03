package main

import (
	"github.com/charmbracelet/log"
	"panzer_server/connect"
	"panzer_server/game"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Panzer Server is starting...")
	game.InitGame()
	connect.Listen()
}
