class Solution {
    public int pivotIndex(int[] nums) {
         int leftSum = 0;
        int totalSum = Arrays.stream(nums).sum();
        for (int i = 0; i < nums.length; leftSum += nums[i++]) {
            if (leftSum == (totalSum - leftSum - nums[i])) {
                return i;
            }
        }

        return -1;
    }
}