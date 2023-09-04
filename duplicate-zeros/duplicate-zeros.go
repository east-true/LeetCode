func duplicateZeros(arr []int)  {
    res := make([]int, 0)
    for i := range arr {
		if arr[i] == 0 {
			res = append(res, 0)
		}

		res = append(res, arr[i])
	}
    
    copy(arr, res)
}