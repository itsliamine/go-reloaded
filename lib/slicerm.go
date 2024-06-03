package lib

func RemoveSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
