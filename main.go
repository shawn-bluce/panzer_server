package main

import (
	"github.com/charmbracelet/log"
	"panzer_server/connect"
)

func main() {
	log.SetLevel(log.DebugLevel)

	log.Debug("Cookie 🍪")
	log.Info("Hello World!")

	connect.Listen()
}
