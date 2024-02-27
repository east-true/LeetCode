func heightChecker(heights []int) int {
    cnt := 0
    exprected := sort(heights)
    for i, exp := range exprected {
        if exp != heights[i] {
            cnt++
        }
    }
    
    return cnt
}

func sort(heights []int) []int {
    copied := make([]int, len(heights))
    copy(copied, heights)
    for i := 0; i < len(copied); i++ {
        for j := i+1; j < len(copied); j++ {
            if copied[i] > copied[j] {
                k := copied[i]
                copied[i] = copied[j]
                copied[j] = k
            }
        }
    }
    return copied
}