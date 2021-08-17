package validate

//IsEmptyByte check if byte data is empty
func IsEmptyByte(data []byte) bool {
	for _, v := range data {
		if v != 0 {
			return false
		}
	}
	return true
}
