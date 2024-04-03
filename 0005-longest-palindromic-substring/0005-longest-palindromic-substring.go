func IsPalindrome(s string) string {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return "-1"
		}
	}

	return s
}

func longestPalindrome(s string) string {
	words := []string{}
	for i := 0; i < len(s); i++ {
		target := s[i]
		words = append(words, string(target))
		for j := i + 1; j < len(s); j++ {
			if target == s[j] {
				word := IsPalindrome(s[i : j+1])
				if word != "-1" {
					words = append(words, word)
				}
			}
		}
	}

	longest := ""
	for _, word := range words {
		if len(longest) < len(word) {
			longest = word
		}
	}

	return longest
}
