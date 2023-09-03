package main

import (
	"github.com/charmbracelet/log"
	"panzer_server/connect"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Info("Panzer Server is starting...")

	connect.Listen()
}
