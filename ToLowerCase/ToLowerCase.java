package leetcode.algorithms;

/**
 * Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.
 */
public class ToLowerCase {
    public String toLowerCase(String str) {
        char[] chars = str.toCharArray();
        for (int i = 0; i < chars.length; i++) {
            int num = (int)chars[i];
            if (num >= 65 && num <= 90) {
                chars[i] = (char) (num + 32);
            }
        }
        return new String(chars);
    }
}
