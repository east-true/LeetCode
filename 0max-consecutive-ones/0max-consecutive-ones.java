class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int cnt = 0;
        int k = 0;
        for (int num : nums) {
            if (num == 0) {
                if (cnt < k) {
                    cnt = k;
                }
                k = 0;
            } else {
                k++;
            }
        }

        return Math.max(cnt, k);
    }
}