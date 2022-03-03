package utils

func ConvertRunesToString(arr []rune) string {
	str := ""

	for _, c := range arr {
		str += string(c)
	}

	return str
}
