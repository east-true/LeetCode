func sortColors(nums []int)  {
    for i := 0; i < len(nums); i++ {
        minIdx := i
        for j := i+1; j < len(nums); j++ {
            if nums[minIdx] > nums[j] {
                minIdx = j
            }
        }
        
        if minIdx != i {
            swap := nums[minIdx]
            nums[minIdx] = nums[i]
            nums[i] = swap
        }
    }
}