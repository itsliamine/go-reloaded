package lib

func BinConvert(s string) string {
	result := 0
	for i := 0; i < len(s); i++ {
		result += int(s[i]) * RecursivePower(2, i)
	}
	return Itoa(result)
}
