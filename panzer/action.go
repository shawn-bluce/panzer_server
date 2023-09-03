package panzer

func UnderAttack(panzer Panzer, attack int) {
	panzer.currentHP -= attack
	if panzer.currentHP <= 0 {
		panzer.status = dead
	}
}

func Shot(panzer Panzer, xSpeed int, ySpeed int) {
	if IsDead(panzer) {
		return
	}
}
