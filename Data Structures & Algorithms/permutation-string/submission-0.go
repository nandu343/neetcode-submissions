func checkInclusion(s1 string, s2 string) bool {
	var targetWindow [26]int
	var window [26]int

	if len(s2) < len(s1) {
		return false
	}

	for _, v := range s1 {
		targetWindow[v - 'a']++
	}

	for i := 0; i < len(s1); i++ {
		window[s2[i] - 'a']++
	}

	if targetWindow == window {
		return true
	}

	l, r := 0, len(s1)

	for r < len(s2) {
		window[s2[r] - 'a']++
		window[s2[l] - 'a']--
		if targetWindow == window {
			return true
		}
		r++
		l++
	}
	return false
}
