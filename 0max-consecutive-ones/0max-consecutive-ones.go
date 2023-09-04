func findMaxConsecutiveOnes(nums []int) int {
    cnt := 0
    k := 0
    for  _, n := range nums {
        if (n == 1) {
            cnt++
        } else {
            if k < cnt {
                k = cnt
            }
            cnt = 0
        }
	}
		
    if k > cnt {
        return k
    } else {
       return cnt
    }
}