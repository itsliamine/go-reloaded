package lib

func GetPrefix(s string) bool {
	for i, char := range s {
		if char != ';' && char != ':' && char != '.' && char != ',' && char != '?' && char != '!' && i != 0 {
			firstCharacter := s[0]
			return !(firstCharacter != 46 && firstCharacter != 58 && firstCharacter != 59 && firstCharacter != 44 && firstCharacter != 63 && firstCharacter != 33)
		}
	}
	return false
}
