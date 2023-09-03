package game

import "panzer_server/panzer"

var panzerList []*panzer.Panzer
var bulletList []*panzer.Bullet

func PanzerJoin(panzer *panzer.Panzer) {
	panzerList = append(panzerList, panzer)
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
