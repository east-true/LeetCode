func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				k := nums[j]
				nums[j] = nums[i]
				nums[i] = k
			}
		}
	}

	half := len(nums) / 2
	median := float64(nums[half])
	if len(nums)%2 == 0 {
		median += float64(nums[half-1])
		median /= 2
	}

	return median
}