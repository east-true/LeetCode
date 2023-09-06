func sortArrayByParity(nums []int) []int {
    arr := make([]int, len(nums))
	copy(arr, nums)

	evenPtr := 0
	oddPtr := len(arr) - 1
	for _, num := range nums {
		if num%2 == 0 {
			arr[evenPtr] = num
			evenPtr++
		} else {
			arr[oddPtr] = num
			oddPtr--
		}
	}

	return arr
}