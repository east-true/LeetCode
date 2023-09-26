class Solution {
    public List<Integer> getRow(int rowIndex) {
        List<Integer> leaf = new ArrayList<>();
        for (int i = 0; i < rowIndex+1; i++) {
            int prev = 1;
            leaf.add(1);
            for (int j = 1; i > 1 && j < i;j++) {
                int curr = leaf.get(j);
                leaf.set(j, prev + curr);
                prev = curr;
            }
        }

        return leaf;
    }
}