// helper to check if a byte is a lowercase letter, uppercase letter, or number
func isAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// helper to convert uppercase characters to lowercase
func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32 // In ASCII, adding 32 turns uppercase into lowercase
	}
	return b
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		// 1. Skip non-alphanumeric characters from the left
		for l < r && !isAlphanumeric(s[l]) {
			l++
		}
		// 2. Skip non-alphanumeric characters from the right
		for l < r && !isAlphanumeric(s[r]) {
			r--
		}

		// 3. Compare the characters ignoring case
		if toLower(s[l]) != toLower(s[r]) {
			return false
		}

		l++
		r--
	}
	return true
}
