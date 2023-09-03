package panzer

const (
	alive = iota
	dead
	ready
	star
)

type Panzer struct {
	// attributes
	maxHP          int
	attack         int
	moveSpeed      int
	bulletSpeed    int
	maxBulletCount int
	shotInterval   int // timestamp by microseconds

	// status
	status             int // alive, dead, ready, star
	currentHP          int
	currentBulletCount int
	nextAllowShotTime  int // timestamp by microseconds
}

func NewPanzer() *Panzer {
	return &Panzer{}
}

func IsDead(panzer Panzer) bool {
	return panzer.status == dead
}

func IsReady(panzer Panzer) bool {
	return panzer.status == ready
}

func IsAlive(panzer Panzer) bool {
	return panzer.status == alive
}

func IsInvulnerable(panzer Panzer) bool {
	return panzer.status == star
}
