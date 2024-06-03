package lib

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
