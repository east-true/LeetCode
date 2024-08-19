func isValid(s string) bool {
	var charArr []rune
	for _, char := range s {
		switch char {
		case '(':
			charArr = append(charArr, ')')
		case '{':
			charArr = append(charArr, '}')
		case '[':
			charArr = append(charArr, ']')
		default:
			if len(charArr) < 1 || charArr[len(charArr)-1] != char {
				return false
			}

			charArr = charArr[:len(charArr)-1]
		}
	}

	return len(charArr) == 0
}