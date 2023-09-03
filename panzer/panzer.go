package panzer

import googleUUID "github.com/google/uuid"

const (
	wait = iota
	ready
	alive
	dead
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
	MaxHP          int
	Attack         int
	MoveSpeed      int
	BulletSpeed    int
	MaxBulletCount int
	ShotInterval   int // timestamp by microseconds

	// status
	Status             int
	Direction          int
	CurrentHP          int
	CurrentBulletCount int
	NextAllowShotTime  int // timestamp by microseconds

	// position
	X, Y int
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
		MaxHP:              maxHP,
		Attack:             attack,
		MoveSpeed:          moveSpeed,
		BulletSpeed:        bulletSpeed,
		MaxBulletCount:     maxBulletCount,
		ShotInterval:       shotInterval,
		Status:             wait,
		Direction:          north,
		CurrentHP:          maxHP,
		CurrentBulletCount: maxBulletCount,
		NextAllowShotTime:  0,
		X:                  0,
		Y:                  0,
	}
}

func IsDead(panzer *Panzer) bool {
	return panzer.Status == dead
}

func IsReady(panzer *Panzer) bool {
	return panzer.Status == ready
}

func IsAlive(panzer *Panzer) bool {
	return panzer.Status == alive
}
