package connect

import (
	"github.com/charmbracelet/log"
	"panzer_server/game"
	panzer "panzer_server/panzer"
)

func processRequest(from, jsonData string) []byte {
	requestData := GetReceivedData(jsonData)

	resp := []byte("Invalid Action")
	if requestData.Action == "join" {
		panzerObj := panzer.NewPanzer(requestData.Name, requestData.UUID, 1, 1, 1, 1, 1, 1)
		game.PanzerJoin(panzerObj)
		log.Info("New Panzer Join", "name", panzerObj.Name, "panzer_list", game.GetPanzerListNames())
		resp = []byte("this is join action")
	} else {
		// if not join, should have uuid
		pz := game.GetPanzerWithUUID(requestData.UUID)
		if pz == nil {
			log.Warn("Not Found Panzer", "uuid", requestData.UUID)
			return resp
		}
	}

	if requestData.Action == "ready" {
		resp = []byte("this is ready action")
	}
	if requestData.Action == "shot" {
		resp = []byte("this is shot action")
	}
	if requestData.Action == "move" {
		resp = []byte("this is move action")
	}
	if requestData.Action == "turn" {
		resp = []byte("this is turn action")
	}
	if requestData.Action == "quit" {
		resp = []byte("this is quit action")
	}
	return resp
}
