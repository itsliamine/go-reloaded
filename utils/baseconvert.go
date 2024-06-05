package lib

func BinConvert(s string) string {
	result := 0
	power := 0
	for i := len(s) - 1; i >= 0; i-- {
		result += int(s[i]-48) * RecursivePower(2, power)
		power++
	}
	return Itoa(result)
}

func HexConvert(s string) string {
	result := 0
	power := 0
	for i := len(s) - 1; i >= 0; i-- {
		if IsLetter(s[i]) {
			if IsLower(s[i]) {
				result += int(s[i]-87) * RecursivePower(16, power)
			} else {
				result += int(s[i]-38) * RecursivePower(16, power)
			}
		} else {
			result += int(s[i]-48) * (RecursivePower(16, power))
		}
		power++
	}
	return Itoa(result)
}
