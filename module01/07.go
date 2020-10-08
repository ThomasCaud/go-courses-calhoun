package module01

func BaseToBase(number string, base int, targetBase int) string {
	return DecToBase(BaseToDec(number, base), targetBase)
}