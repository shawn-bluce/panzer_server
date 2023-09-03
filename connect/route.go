package connect

import (
	"github.com/charmbracelet/log"
	"panzer_server/game"
	panzer "panzer_server/panzer"
)

func processRequest(from, jsonData string) []byte {
	requestData := GetReceivedData(jsonData)

	resp := []byte("{\"message\": \"Operation Failed\"}")
	if requestData.Action == "join" {
		panzerObj := panzer.NewPanzer(requestData.Name, requestData.UUID, 1, 1, 1, 1, 1, 1)
		success := game.PanzerJoin(panzerObj)
		if !success {
			return []byte("{\"message\": \"panzer list is full\"}")
		}
		resp = GenerateResponseData(panzerObj.UUID)
	} else {
		// if not join, should have uuid
		pz := game.GetPanzerWithUUID(requestData.UUID)
		if pz == nil {
			log.Warn("Not Found Panzer", "uuid", requestData.UUID)
			return resp
		}
	}

	if requestData.Action == "ready" {
		resp = []byte("{}")
	}
	if requestData.Action == "shot" {
		resp = []byte("{}")
	}
	if requestData.Action == "move" {
		resp = []byte("{}")
	}
	if requestData.Action == "turn" {
		resp = []byte("{}")
	}
	if requestData.Action == "quit" {
		resp = []byte("{}")
	}
	return resp
}
