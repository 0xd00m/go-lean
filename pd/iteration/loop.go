package loop

func loop(x string, y int) string {
	var z string
	for i := 1; i < y; i++ {
		z += x
	}
	return z
}
