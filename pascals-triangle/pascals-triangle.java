class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> tree = new ArrayList<>();
            List<Integer> leaf = new ArrayList<>();
            int prev = 0;
            for (int i = 0; i < numRows; i++) {
                leaf.add(1);
                prev = leaf.get(0);
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