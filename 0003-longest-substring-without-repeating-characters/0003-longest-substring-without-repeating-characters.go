func lengthOfLongestSubstring(s string) int {
	subarr := []string{}
	for s != "" {
		n, sub := GetSubstring(s)
		if n != 0 {
			subarr = append(subarr, sub)
			s = s[1:]
		}
	}

	max := 0
	for _, sub := range subarr {
		if max < len(sub) {
			max = len(sub)
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
