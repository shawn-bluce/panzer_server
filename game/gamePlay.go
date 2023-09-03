package game

import (
	"github.com/charmbracelet/log"
	"panzer_server/panzer"
)

var panzerList []*panzer.Panzer
var bulletList []*panzer.Bullet
var gameStatus int

const (
	waiting = iota
	gaming
)

func InitGame() {
	panzerList = []*panzer.Panzer{}
	bulletList = []*panzer.Bullet{}
	gameStatus = waiting
	log.Info("Init Game, create blank panzerList and bulletList, set gameStatus to waiting")
}

func StartGame() bool {
	if AllIsReady() {
		log.Info("All panzer are ready, start game")
		gameStatus = gaming
		return true
	}
	log.Warn("Not all panzer are ready, start game failed")
	return false
}

func EndGame() {
	log.Info("End game, set gameStatus to waiting")
	gameStatus = waiting
}

func PanzerJoin(panzer *panzer.Panzer) bool {
	if len(panzerList) >= 4 {
		log.Warn("Panzer join failed, because the panzer list is full")
		return false
	}
	panzerList = append(panzerList, panzer)
	log.Info("New Panzer Join", "Name", panzer.Name)
	return true
}

func GetPanzerListNames() string {
	resp := ""
	for _, p := range panzerList {
		resp += p.Name + ", "
	}
	return resp
}

func GetPanzerListExcludedUUID(uuid string) []*panzer.Panzer {
	resp := []*panzer.Panzer{}
	for _, p := range panzerList {
		if p.UUID != uuid {
			resp = append(resp, p)
		}
	}
	return resp
}

func GetPanzerWithUUID(uuid string) *panzer.Panzer {
	for _, p := range panzerList {
		if p.UUID == uuid {
			return p
		}
	}
	return nil
}

func AllIsReady() bool {
	for _, p := range panzerList {
		if panzer.IsReady(p) {
			return false
		}
	}
	return true
}

func BulletJoin(bullet *panzer.Bullet) {
	bulletList = append(bulletList, bullet)
}
