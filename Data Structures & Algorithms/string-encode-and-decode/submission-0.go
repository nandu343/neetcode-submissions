type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var ans string
	for _, v := range strs {
		// Add a delimiter (like '#') so we know where the length ends
		ans += strconv.Itoa(len(v)) + "#" + v
	}
	return ans
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	i := 0

	// Process the entire encoded string
	for i < len(encoded) {
		j := i
		// Find where the '#' delimiter is to read the length
		for encoded[j] != '#' {
			j++
		}

		// Extract the length and convert it to an integer
		length, _ := strconv.Atoi(encoded[i:j])
		
		// Move our pointer past the '#'
		i = j + 1
		
		// Extract the actual string using the length
		res = append(res, encoded[i:i+length])
		
		// Move the pointer to the start of the next encoded block
		i += length
	}

	return res
}
