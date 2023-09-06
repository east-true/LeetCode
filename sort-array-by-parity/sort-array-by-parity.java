class Solution {
    public int[] sortArrayByParity(int[] nums) {
        int[] arr = new int[nums.length];
        int evenPtr = 0;
        int oddPtr = arr.length-1;
        for (int num : nums) {
            if (num % 2 == 0) {
                arr[evenPtr++] = num;
            } else {
                arr[oddPtr--] = num;
            }
        }

        return arr;
    }
}