

            
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func checkPalindrome(inputString string) bool {
	rev:=Reverse(inputString)
	for k, i := range inputString {
		if rev[k] != byte(i) {
			return false
		}
	}

	return true
}
