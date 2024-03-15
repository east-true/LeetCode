func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    return x == Reverse(x)
}

func Reverse(num int) int {
	result := 0
	for num > 0 {
		result = (result * 10) + (num % 10)
		num /= 10
	}

	return result
}