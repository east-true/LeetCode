func getRow(rowIndex int) []int {
    leaf := make([]int, 0)
	for i := 0; i < rowIndex+1; i++ {
		prev := 1
		leaf = append(leaf, 1)
		for j := 1; i > 1 && j < i; j++ {
			curr := leaf[j]
			leaf[j] += prev
			prev = curr
		}
	}

	return leaf
}