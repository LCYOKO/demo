package algo

func isPalindrome(x int) bool {
	str := string(rune(x))
	for i, j := 0, len(str); i < j; {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
