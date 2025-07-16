package function

func CheckRange(inputText string) bool {
	for _, v := range inputText{
		if v < 32 || v > 126 {
			return false
		}
	}
	return true
}