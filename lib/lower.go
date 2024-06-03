package lib

func ToLower(s string) string {
	sb := []byte(s)
	for i, l := range sb {
		if IsLetter(l) && !IsLower(l) {
			sb[i] = l + 32
		}
	}
	return string(sb)
}
