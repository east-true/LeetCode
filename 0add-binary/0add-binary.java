class Solution {
    public String addBinary(String a, String b) {
        boolean up = false;
        StringBuilder sb = new StringBuilder();

        // a,b length
        int gap = a.length() - b.length();
        if (gap > 0) {
            for (int i = 0; i < gap; i++) {
                b = "0" + b;
            }
        } else if (gap < 0) {
            for (int i = 0; i > gap; i--) {
                a = "0" + a;
            }
        }

        for (int i = a.length()-1; i >= 0; i--) {
            if (a.charAt(i) != b.charAt(i)) {
                if (up) {
                    sb.append('0');
                } else {
                    sb.append('1');
                }
            } else {
                if (a.charAt(i) == '1' && b.charAt(i) == '1') {
                    
                    if (up) {
                        sb.append('1');
                    } else {
                        sb.append('0');
                    }
                    up = true;
                } else {
                    if (up) {
                        sb.append('1');
                        up = false;
                    } else {
                        sb.append('0');
                    }
                }
            }
        }

        if (up) {
            sb.append(1);
        }

        return sb.reverse().toString();
    }
}