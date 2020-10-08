package module01

func Reverse(chaine string) string {
	res := ""

	for _, val := range chaine {
		res = string(val) + res
	}

	return res
}