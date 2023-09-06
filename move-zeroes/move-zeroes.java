class Solution {
    public void moveZeroes(int[] nums) {
        int ptr = 0;
        int store;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != 0) {
                store = nums[ptr];
                nums[ptr++] = nums[i];
                nums[i] = store;
            }
        }
    }
}