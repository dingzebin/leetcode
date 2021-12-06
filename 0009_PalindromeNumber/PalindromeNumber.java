package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 9. Palindrome Number
 * @createTime 2020/10/25 21:55
 */
public class PalindromeNumber {
    public static void main(String[] args) {
        System.out.println(new PalindromeNumber().isPalindrome(121));
    }
    public boolean isPalindrome(int x) {
        int tmp = x;
        if (tmp < 0) return false;
        int result = 0;
        while (tmp != 0) {
            int num = tmp % 10;
            // judge whether overflow
            if (result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE / 10 && num > Integer.MAX_VALUE % 10)) {
                return false;
            }
            result = result * 10 + num;
            tmp /= 10;
        }
        return result == x ? true : false;
    }
}
