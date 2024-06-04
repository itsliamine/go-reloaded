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

func ToUpper(s string) string {
	sb := []byte(s)
	for i, l := range sb {
		if IsLetter(l) && !IsUpper(l) {
			sb[i] = l - 32
		}
	}
	return string(sb)
}

func IsLetter(l byte) bool {
	if l >= 65 && l <= 90 || l >= 97 && l <= 122 {
		return true
	}
	return false
}

func IsUpper(l byte) bool {
	if l >= 65 && l <= 90 {
		return true
	}
	return false
}

func IsLower(l byte) bool {
	if l >= 97 && l <= 122 {
		return true
	}
	return false
}

func IsVowel(l byte) bool {
	if l == 65 || l == 97 || l == 69 || l == 101 || l == 73 || l == 105 || l == 79 || l == 111 || l == 85 || l == 117 || l == 89 || l == 121 || l == 72 || l == 104 {
		return true
	}
	return false
}
