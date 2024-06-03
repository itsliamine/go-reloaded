package lib

func IsLetter(l byte) bool {
	if l >= 65 && l <= 90 || l >= 97 && l <= 122 {
		return true
	}
	return false
}

func IsUpperL(l byte) bool {
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
