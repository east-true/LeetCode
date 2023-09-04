class Solution {
    public int findNumbers(int[] nums) {
        int cnt = 0;
        for (int num : nums) {
            String numeric = Integer.toString(num);
            if ((numeric.length() % 2) == 0) {
                cnt++;
            }
        }

        return cnt;
    }
}