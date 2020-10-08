package module01

func Sum(list []int) int {
	total := 0

	for _, val := range list {
		total += val
	}

	return total
}