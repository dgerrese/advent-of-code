package util

func IntLen(i int) int {
	if i <= 0 {
		return 0
	}

	return 1 + IntLen(i/10)
}
