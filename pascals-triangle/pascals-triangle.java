class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> tree = new ArrayList<>();
        List<Integer> leaf = new ArrayList<>();
        for (int i = 0; i < numRows; i++) {
            int prev = 1;
            leaf.add(1);
            for (int j = 1; i > 1 && j < i;j++) {
                int curr = leaf.get(j);
                leaf.set(j, prev + curr);
                prev = curr;
            }

            tree.add(new ArrayList<>(leaf));
        }

        return tree;
    }
}