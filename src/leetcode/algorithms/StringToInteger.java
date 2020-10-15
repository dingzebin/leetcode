package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 8. String to Integer
 * @createTime 2020/10/15 23:27
 */
public class StringToInteger {

    public static void main(String[] args) {
        System.out.println(new StringToInteger().myAtoi("21474836460"));
    }
    // max int: 2147483647  min int: -2147483648
    public int myAtoi(String s) {
        int i = 0;
        while (i < s.length() && s.charAt(i) == ' ') {
            i++;
        }
        // empty string
        if (i >= s.length()) return 0;
        int sign = 1;
        if (s.charAt(i) == '-') {
            sign = -1;
            i++;
        } else if (s.charAt(i) == '+') {
            i++;
        }
        int result = 0;
        while (i < s.length() && s.charAt(i) >= '0' && s.charAt(i) <= '9') {
            if (result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE / 10 && s.charAt(i) - '0' > Integer.MAX_VALUE % 10)) {
                return sign == -1 ? Integer.MIN_VALUE : Integer.MAX_VALUE;
            }
            result = result * 10 + s.charAt(i++) - '0';
        }
        return result * sign;
    }
}
