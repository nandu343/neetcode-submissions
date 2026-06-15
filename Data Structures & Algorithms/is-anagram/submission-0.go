func isAnagram(s string, t string) bool {
	var ana [26]int
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		ana[s[i] - 'a']++
		ana[t[i]-'a']--
	}

	for _, v := range ana {
		if v != 0 {
			return false
		}
	}
	return true
}
