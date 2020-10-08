package module01

func FindTwoThatSum(numbers []int, sum int) (int, int) {
	for i, valA := range numbers {
		for j, valB := range numbers {
			if i == j {
				continue
			}
			if valA + valB == sum {
				return i, j
			}
		}
	}
	return -1, -1
}