package player

import "github.com/google/uuid"
import panzer "panzer_server/panzer"

var PlayerArray = make([]*Player, 0)

type Player struct {
	name   string
	panzer *panzer.Panzer
	uuid   string
}

func NewPlayer(name string) *Player {
	return &Player{
		name:   name,
		panzer: panzer.NewPanzer(),
		uuid:   uuid.New().String(),
	}
}

func PlayerJoin(player *Player) {
	PlayerArray = append(PlayerArray, player)
}

func PlayerLeave(player *Player) {
	for i, p := range PlayerArray {
		if p == player {
			PlayerArray = append(PlayerArray[:i], PlayerArray[i+1:]...)
			return
		}
	}
}
