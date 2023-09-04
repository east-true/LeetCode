func findNumbers(nums []int) int {
    cnt := 0
    for _, n := range nums {
        fmt.Println(len(string(n)))
        if (len(strconv.Itoa(n)) % 2) == 0 {
            cnt++
        }
    }
    
    return cnt
}