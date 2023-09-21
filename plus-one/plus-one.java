class Solution {
    public int[] plusOne(int[] digits) {
        digits[digits.length - 1]++;
        if (digits[digits.length-1] == 10) {
            for (int i = digits.length - 1; i > 0; i--) {
                if (digits[i] == 10) {
                    digits[i - 1] += 1;
                    digits[i] = 0;
                }
            }

            if (digits[0] == 10) {
                int[] newDigits = new int[digits.length+1];
                newDigits[0] = 1;
                newDigits[1] = 0;
                if (digits.length - 2 >= 0) System.arraycopy(digits, 2, newDigits, 3, digits.length - 2);
                return newDigits;
            }
        }
        return digits;
    }
}