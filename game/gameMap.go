package game

type GameMap struct {
	xMin, xMax, yMin, yMax int
}

func InitGameMap(x, y int) GameMap {
	return GameMap{
		xMin: 0,
		xMax: x,
		yMin: 0,
		yMax: y,
	}
}
