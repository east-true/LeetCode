func reverse(x int) int {
    res := 0
	for x != 0 {
		res = (res + (x % 10)) * 10
		x = x / 10
	}

	res = res / 10
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}