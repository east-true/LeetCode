func removeElement(nums []int, val int) int {
    cnt := 0
    for i, n := range nums {
        if n == val {
            nums[i] = 51
        } else {
            cnt++
        }
    }
    
    sort.Ints(nums)
    return cnt
}