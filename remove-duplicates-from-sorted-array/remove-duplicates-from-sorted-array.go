func removeDuplicates(nums []int) int {
    ptr := 1
    for i := 1; i < len(nums); i++ {
        if (nums[i-1] != nums[i]) {
            nums[ptr] = nums[i]
            ptr++
        }
    }
    
    return ptr
}