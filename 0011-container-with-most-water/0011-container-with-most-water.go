func maxArea(height []int) int {
    var i, j, water, area int
	j = len(height) - 1
	for i < j {
		water = (j - i) * min(height[i], height[j])
		area = max(area, water)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return area
}