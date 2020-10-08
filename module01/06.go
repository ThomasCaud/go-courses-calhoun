package module01

func BaseToDec(val string, base int) int {
	const charset = "0123456789ABCDEF"
	res := 0
	multiplier := 1

	for i := len(val) - 1 ; i >= 0 ; i-- {
		index := -1

		for j, char := range charset {
			if char == rune(val[i]) {
				index = j
			}
		}

		res += index * multiplier
		multiplier *= base
	}

	return res
}