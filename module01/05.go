package module01

func DecToBase(dec int, base int) string {
	var charset = "0123456789ABCDEF"
	var res string

	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec /= base
	}

	return res
}