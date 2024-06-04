package lib

func Contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	var arr []string
	temp := ""
	for i, l := range s {
		if l == ' ' && i != 0 {
			if i < len(s)-1 {
				if s[i+1] != ' ' {
					arr = append(arr, temp)
					temp = ""
				}
			}
		} else if l != ' ' {
			temp += string(l)
		}
		if i == len(s)-1 {
			arr = append(arr, temp)
		}
	}
	return arr
}

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

func Index(s string, toFind string) int {
	n := len(toFind)
	if n == 0 {
		return 0
	} else if n == 1 {
		for i, l := range s {
			if l == []rune(toFind)[0] {
				return i
			}
		}
	} else if n > len(s) {
		return -1
	} else if n == len(s) {
		if toFind == s {
			return 0
		}
		return -1
	}
	for i := 0; i < len(s); i++ {
		if i == len(s)-len(toFind) {
			return -1
		}
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return -1
}
