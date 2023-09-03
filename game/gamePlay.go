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

func BulletJoin(bullet *panzer.Bullet) {
	bulletList = append(bulletList, bullet)
}
