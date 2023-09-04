func sortedSquares(nums []int) []int {
    for i := range nums {
		if nums[i] < 0 {
			nums[i] *= -1
		}
        
        nums[i] *= nums[i]
	}
    
    sort.Ints(nums)
    return nums
}