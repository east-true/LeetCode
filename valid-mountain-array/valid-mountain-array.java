class Solution {
    public boolean validMountainArray(int[] arr) {
        int top = 0;
        for (int i = 1; i < arr.length; i++) {
            if (arr[i-1] >= arr[i]) {
                break;
            }

            top = i;
        }

        if (top == 0 || top == arr.length-1) {
            return false;
        }

        for (int i = top+1; i < arr.length; i++) {
            if (arr[i] >= arr[i-1]) {
                return false;
            }
        }

        return true;
    }
}