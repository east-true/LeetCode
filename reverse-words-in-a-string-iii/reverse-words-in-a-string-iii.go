func reverseWords(s string) string {
	bs := bytes.NewBufferString("")
	words := strings.Split(s, " ")

	for _, word := range words {
		for j := len(word) - 1; j >= 0; j-- {
			bs.WriteByte(word[j])
		}
		bs.WriteString(" ")
	}

    return strings.TrimSpace(bs.String())
}