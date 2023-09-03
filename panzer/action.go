package panzer

func UnderAttack(pz *Panzer, attack int) {
	pz.CurrentHP -= attack
	if pz.CurrentHP <= 0 {
		pz.Status = dead
	}
}

func Shot(pz *Panzer, xSpeed int, ySpeed int) {
	if IsDead(pz) {
		return
	}
}
