func validMountainArray(arr []int) bool {
    top := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			break
		}

		top = i
	}

	if top == 0 || top == len(arr)-1 {
		return false
	}

	for i := top + 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			return false
		}
	}

	return true
}