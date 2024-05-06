
func Hero(bullets, dragons int) bool {
	b := bullets
	d := dragons
	x := d * 2
	if b >= x {
		return true
	}
	return false

}

func Hero(bullets, dragons int) bool {
	b := bullets
	d := dragons
	x := d * 2
	if b >= x {
		return true
	}
	return false

}



package kata

func Hero(bullets, dragons int) bool {
  return bullets >= 2*dragons
}