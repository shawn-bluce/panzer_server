package panzer

func UnderAttack(pz *Panzer, attack int) {
	pz.currentHP -= attack
	if pz.currentHP <= 0 {
		pz.status = dead
	}
}

func Shot(pz *Panzer, xSpeed int, ySpeed int) {
	if IsDead(pz) {
		return
	}
}
