package lib

func Capitalize(s string) string {
	h := []rune(s)
	result := ""
	isPreviousAlpha := true
	for i := 0; i <= len(h)-1; i++ {
		if isPreviousAlpha {
			if h[i] >= 'a' && h[i] <= 'z' {
				result += string(h[i] - 32)
			} else {
				result += string(h[i])
			}
		} else {
			if h[i] >= 'A' && h[i] <= 'Z' {
				result += string(h[i] + 32)
			} else {
				result += string(h[i])
			}
		}
		isPreviousAlpha = !((h[i] >= 'a' && h[i] <= 'z') || (h[i] >= 'A' && h[i] <= 'Z') || (h[i] >= '0' && h[i] <= '9'))
	}
	return result
}
