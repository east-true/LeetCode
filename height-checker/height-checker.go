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
    expected := make([]int, len(heights))
    copy(expected, heights)
    for i := 0; i < len(expected); i++ {
        for j := i+1; j < len(expected); j++ {
            if expected[i] > expected[j] {
                k := expected[i]
                expected[i] = expected[j]
                expected[j] = k
            }
        }
    }
    return expected
}