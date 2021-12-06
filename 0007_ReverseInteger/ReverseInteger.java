package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 7. Reverse Integer
 * @createTime 2020/10/25 21:29
 */
public class ReverseInteger {
    public static void main(String[] args) {
        System.out.println(new ReverseInteger().reverse(-214748364));
    }
    // max int: 2147483647  min int: -2147483648
    public int reverse(int x) {
        int result = 0;
        while (x != 0) {
            int num = x % 10;
            // judge whether overflow
            if (result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE / 10 && num > Integer.MAX_VALUE % 10)
                || result < Integer.MIN_VALUE / 10 || (result == Integer.MIN_VALUE / 10 && num < Integer.MIN_VALUE % 10)) {
                return 0;
            }
            result = result * 10 + num;
            x /= 10;
        }
        return result;
    }
}
