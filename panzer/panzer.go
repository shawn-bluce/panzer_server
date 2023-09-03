package panzer

import googleUUID "github.com/google/uuid"

const (
	wait = iota
	ready
	alive
	dead
	star
)

const (
	north = iota
	south
	west
	east
)

type Panzer struct {
	// base
	Name string
	UUID string

	// attributes
	maxHP          int
	attack         int
	moveSpeed      int
	bulletSpeed    int
	maxBulletCount int
	shotInterval   int // timestamp by microseconds

	// status
	status             int
	direction          int
	currentHP          int
	currentBulletCount int
	nextAllowShotTime  int // timestamp by microseconds

	// position
	x, y int
}

func NewPanzer(name, uuid string, maxHP, attack, moveSpeed, bulletSpeed, maxBulletCount, shotInterval int) *Panzer {
	uuid = googleUUID.New().String()

	// only one Panzer type now.
	maxHP = 100
	attack = 20
	moveSpeed = 20
	bulletSpeed = 40
	maxBulletCount = 10
	shotInterval = 500 // microsecond

	return &Panzer{
		Name:               name,
		UUID:               uuid,
		maxHP:              maxHP,
		attack:             attack,
		moveSpeed:          moveSpeed,
		bulletSpeed:        bulletSpeed,
		maxBulletCount:     maxBulletCount,
		shotInterval:       shotInterval,
		status:             wait,
		direction:          north,
		currentHP:          maxHP,
		currentBulletCount: maxBulletCount,
		nextAllowShotTime:  0,
		x:                  0,
		y:                  0,
	}
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
