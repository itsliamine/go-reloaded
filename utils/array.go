package lib

func Insert(a []string, index int, value string) []string {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func RemoveSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
