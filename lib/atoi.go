package lib

func Atoi(s string) int {
	result := 0
	sign := 1
	if len(s) > 0 {
		if s[0] == 45 {
			sign = -1
		}
		for i, e := range s {
			if i == 0 && int(e) == 45 {
				sign = -1
			} else if i == 0 && int(e) == 43 {
				sign = 1
			} else if int(e) < 48 || int(e) > 57 && i != 0 {
				return 0
			} else if int(e) != 45 && int(e) != 43 {
				result = 10*result + (int(e) - 48)
			}
		}
	}
	return result * sign
}
