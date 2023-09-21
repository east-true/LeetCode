class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> list = new ArrayList<>();
        int left = 0;
        int top = 0;
        int right = matrix[0].length-1;
        int down = matrix.length-1;
        while (left <= right && top <= down) {
            //right
            for (int i = left; i <= right; i++) {
                list.add(matrix[top][i]);
            }
            top++;

            // down
            for (int i = top; i <= down; i++) {
                list.add(matrix[i][right]);
            }
            right--;

            // left
            if (top <= down) {
                for (int i = right; i >= left; i--) {
                    list.add(matrix[down][i]);
                }
            }
            down--;

            // top
            if (left <= right) {
                for (int i = down; i >= top; i--) {
                    list.add(matrix[i][left]);
                }
            }
            left++;
        }

        return list;
    }
}