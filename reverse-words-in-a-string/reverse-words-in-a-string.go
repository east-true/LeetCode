func reverseWords(s string) string {
    words := strings.Split(s, " ")
	reverse := make([]string, 0)
	ptr := 0
	for i := len(words) - 1; i >= 0; i-- {
		word := strings.TrimSpace(words[i])
		if word != "" {
			reverse = append(reverse, word)
			ptr++
		}
	}

	return strings.Join(reverse, " ")
}
