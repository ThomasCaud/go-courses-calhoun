package module01

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}

func GCD(a, b int) int {
	for i := min(a, b) ; i > 0 ; i-- {
		if a % i == 0 && b % i == 0 {
			return i
		}
	}
	return 1
}