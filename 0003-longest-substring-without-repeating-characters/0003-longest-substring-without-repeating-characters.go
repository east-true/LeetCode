func lengthOfLongestSubstring(s string) int {
	max := 0
	for s != "" {
		n, sub := GetSubstring(s)
		if n != 0 {
			if max < len(sub) {
				max = len(sub)
			}
			s = s[1:]
		}
	}

	return max
}

func GetSubstring(s string) (int, string) {
	m := map[rune]byte{}
	for i, char := range s {
		if _, ok := m[char]; ok {
			return i, s[:i]
		} else {
			m[char] = 0
		}
	}

	return len(s), s
}
