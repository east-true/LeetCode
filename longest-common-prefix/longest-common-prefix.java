class Solution {
    public String longestCommonPrefix(String[] strs) {
        StringBuilder prefix = new StringBuilder();
        String longest = "";
        for (int i = 0; i < strs.length; i++) {
            if (longest.length() < strs[i].length()) {
                longest = strs[i];
            }
        }

        for (int i = 0; i < longest.length(); i++) {
            prefix.append(longest.charAt(i));
            for (int j = 0; j < strs.length; j++) {
                if (strs[j].length() > 0 && strs[j].length() >= i) {
                    if (!strs[j].startsWith(prefix.toString())) {
                        prefix.deleteCharAt(prefix.lastIndexOf(String.valueOf(longest.charAt(i))));
                        return prefix.toString();
                    }
                } else {
                    prefix.deleteCharAt(prefix.lastIndexOf(String.valueOf(longest.charAt(i))));
                    break;
                }
            }
        }

        return prefix.toString();
    }
}