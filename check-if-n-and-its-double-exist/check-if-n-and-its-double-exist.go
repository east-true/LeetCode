func checkIfExist(arr []int) bool {
    for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
            if arr[i] == (2 * arr[j]) {
                return true
            } else if arr[j] == (2 * arr[i]) {
                return true
            }
		}
	}

	return false;
}