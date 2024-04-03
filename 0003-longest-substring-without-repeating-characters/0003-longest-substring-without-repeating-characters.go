func lengthOfLongestSubstring(s string) int {
	max := 0
	for s != "" {
		n := GetSubstring(s)
		if n != 0 {
			if max < n {
				max = n
			}
			s = s[1:]
		}
	}

	return max
}

func GetSubstring(s string) int {
	m := map[rune]byte{}
	for i, char := range s {
		if _, ok := m[char]; ok {
			return i
		} else {
			m[char] = 0
		}
	}

	return len(s)
}
