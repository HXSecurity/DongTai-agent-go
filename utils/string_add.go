package utils

func StringAdd(s ...string) string {
	byteArr := []byte("")
	for _, v := range s {
		byteArr = append(byteArr, []byte(v)...)
	}
	return string(byteArr)
}
