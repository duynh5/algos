package commonprefixscan

// LongestCommonPrefix returns the longest byte prefix shared by all strings.
// It returns an empty string when strs is empty.
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for column := 0; column < len(strs[0]); column++ {
		expected := strs[0][column]
		for row := 1; row < len(strs); row++ {
			if column >= len(strs[row]) || strs[row][column] != expected {
				return strs[0][:column]
			}
		}
	}
	return strs[0]
}
