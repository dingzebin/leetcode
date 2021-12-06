package leetcode.algorithms;

/**
 *  Given a string, find the length of the longest substring without repeating characters.
 */
public class LengthOfLongestSubstring {
    public static void main(String[] args) {
      String s = "abcabcbb";
      System.out.println(new LengthOfLongestSubstring().lengthOfLongestSubstring(s));
    }
    public int lengthOfLongestSubstring(String s) {
        int[] arr = new int[128];
        int max = 0;
        for (int i = 0, j = 0;i < s.length(); i++) {
            if (arr[s.charAt(i)] != 0) {
                j = Math.max(j, arr[s.charAt(i)]);
            }
            arr[s.charAt(i)] = i + 1;
            max = Math.max(max, i - j + 1);
        }
        return max;
    }
}
