package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 13. Roman to Integer
 * @createTime 2020/10/25 22:05
 */
public class RomanToInteger {
    public static void main(String[] args) {
        System.out.println(new RomanToInteger().romanToInt("MCMXCIV"));
    }

    public int romanToInt(String s) {
        char[] chars = s.toCharArray();
        int result = 0;
        int lastNum = 0;
        int cur = 0;
        for(int i = 0; i < chars.length; i++) {
            switch (chars[i]) {
                case 'I': cur = 1; break;
                case 'V': cur = 5; break;
                case 'X': cur = 10; break;
                case 'L': cur = 50; break;
                case 'C': cur = 100; break;
                case 'D': cur = 500; break;
                case 'M': cur = 1000; break;
            }
            if (lastNum != 0 && lastNum < cur) {
                result += (cur - 2 * lastNum);
                lastNum = 0;
            } else {
                result += cur;
                lastNum = cur;
            }
        }
        return result;
    }
}
