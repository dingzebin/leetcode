package leetcode.algorithms;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author dzeb
 * @version 1.0
 * @Description Given a string s, return the longest palindromic substring in s.
 * @createTime 2021/5/3 14:46
 */
public class LongestPalindromicSubstring5 {

    public static void main(String[] args) {
        String s = "abac";
        System.out.println(new LongestPalindromicSubstring5().longestPalindrome(s));
    }


    public String longestPalindrome(String s) {
        int len = s.length();
        if (s.length() < 2) return s;

        byte[][] dp = new byte[len][len];
        for (int i = 0; i < len; i++) {
            dp[i][i] = 0;
        }
        int begin = 0;
        int maxLen = 1;
        char[] chars = s.toCharArray();
        Map<Character, List<Integer>> char_index = new HashMap<>();
        for (int i = 0; i < len; i++) {
            char c = chars[i];
            List<Integer> indexs = char_index.get(c);
            if (indexs == null) {
                indexs = new ArrayList<>();
                char_index.put(c, indexs);
            }
            for (Integer index : indexs) {
                byte flag = isPalindrome(chars, index + 1, i - 1, dp);
                if (flag == 1) {
                    if (i - index + 1 > maxLen) {
                        begin = index;
                        maxLen = i - index + 1;
                    }
                    break;
                }
                dp[index][i] = flag;
            }
            indexs.add(i);
        }
        return s.substring(begin, begin + maxLen);
    }

    private byte isPalindrome(char[] chars, int start, int end, byte[][] dp) {
        if (start >= end) {
            return 1;
        }
        if (dp[start][end] != 0) {
            return dp[start][end];
        }
        if (chars[start] == chars[end]) {
            byte flag = isPalindrome(chars, start + 1, end - 1, dp);
            dp[start][end] = flag;
            return flag;
        } else {
            dp[start][end] = -1;
            return -1;
        }
    }
}
