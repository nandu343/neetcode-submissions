func minWindow(original string, check string) string {
    if len(original) == 0 || len(check) == 0 || len(original) < len(check) {
		return ""
	}

	var tarwindow [128]int
	requiredUnique := 0
	for i := 0; i < len(check); i++ {
		if tarwindow[check[i]] == 0 {
			requiredUnique++
		}
		tarwindow[check[i]]++
	}

	var window [128]int
	l := 0
	minLen := math.MaxInt32
	startIdx := 0
	matchedUnique := 0

	for r := 0; r < len(original); r++ {
		charR := original[r]
		window[charR]++

		if tarwindow[charR] > 0 && window[charR] == tarwindow[charR] {
			matchedUnique++
		}

		for matchedUnique == requiredUnique {
			currentWindowLen := r - l + 1
			
			// TIE BREAKER LOGIC:
			// 1. If it's strictly smaller, update it.
			// 2. If it's equal, pick the one that is lexicographically smaller.
			if currentWindowLen < minLen {
				minLen = currentWindowLen
				startIdx = l
			} else if currentWindowLen == minLen {
				currentStr := original[l : l+currentWindowLen]
				bestStr := original[startIdx : startIdx+minLen]
				if currentStr < bestStr {
					startIdx = l
				}
			}

			charL := original[l]
			if tarwindow[charL] > 0 && window[charL] == tarwindow[charL] {
				matchedUnique--
			}
			window[charL]--
			l++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return original[startIdx : startIdx+minLen]
}
