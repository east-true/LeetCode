func generate(numRows int) [][]int {
	tree := make([][]int, 0)
	leaf := make([]int, 0)
	for i := 0; i < numRows; i++ {
		prev := 1
		leaf = append(leaf, 1)
		for j := 1; i > 1 && j < i; j++ {
			curr := leaf[j]
			leaf[j] += prev
			prev = curr
		}

		copyLeaf := make([]int, len(leaf))
		copy(copyLeaf, leaf)
		tree = append(tree, copyLeaf)
	}

	return tree
}