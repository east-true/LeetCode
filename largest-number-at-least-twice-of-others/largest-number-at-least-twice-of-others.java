class Solution {
    public int dominantIndex(int[] nums) {
        int idx = 0;
        for (int i = 1; i < nums.length; i++)
            if(nums[idx] < nums[i])
                idx = i;

        for (int num : nums) {
            if (nums[idx] != num && nums[idx] < (num * 2)) {
                System.out.println(num);
                return -1;
            }
        }

        return idx;
    }
}