package lib

func HexConvert(s string) string {
	result := 0
	for i := 0; i < len(s); i++ {
		if IsLetter(s[i]) {
			result += int(s[i]-38) * RecursivePower(16, i)
		} else {
			result += int(s[i]) * RecursivePower(16, i)
		}
	}
	return Itoa(result)
}
