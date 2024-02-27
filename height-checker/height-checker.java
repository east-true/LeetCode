class Solution {
    public int heightChecker(int[] heights) {
        int cnt = 0;
        int[] expected = heights.clone();
        Arrays.sort(expected);
        for (var i = 0; i < expected.length; i++) {
            if (heights[i] != expected[i]) {
                cnt++;
            }
        }
        
        return cnt;
    }
}