func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = int(uint(nums[i]) * uint(nums[i]))
	}

	sort.Ints(nums)
	return nums
}