package lib

func ToUpper(s string) string {
	sb := []byte(s)
	for i, l := range sb {
		if IsLetter(l) && !IsUpperL(l) {
			sb[i] = l - 32
		}
	}
	return string(sb)
}
