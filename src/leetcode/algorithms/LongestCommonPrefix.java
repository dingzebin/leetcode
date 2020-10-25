package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 14. Longest Common Prefix
 * @createTime 2020/10/25 22:22
 */
public class LongestCommonPrefix {
    public static void main(String[] args) {
        System.out.println(new LongestCommonPrefix().longestCommonPrefix(new String[]{"ab", "a"}));
    }

    public String longestCommonPrefix(String[] strs) {
        if (strs.length == 0) {
            return "";
        }
        if (strs.length == 1) {
            return strs[0];
        }
        StringBuilder sb = new StringBuilder();
        char cur;
        for (int i = 0, len0 = strs[0].length(); i < len0; i++) {
            cur = strs[0].charAt(i);
            for(int j = 1; j < strs.length; j++) {
                if (i >= strs[j].length() || strs[j].charAt(i) != cur) {
                    return sb.toString();
                }
            }
            sb.append(cur);
        }
        return sb.toString();
    }
}
