package module01

func Factor(primes []int, factor int) []int {
	var res []int

	for _, prime := range primes {
		for factor % prime == 0 {
			factor /= prime
			res = append(res, prime)
		} 
	}

	return res;
}