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

func RemoveChar(s string, index int) string {
	a := []rune(s)
	a = append(a[:index], a[index+1:]...)
	return string(a)
}

func Insert(s string, index int, value rune) string {
	a := []rune(s)
	if len(a) == index {
		return string(append(a, value))
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return string(a)
}
