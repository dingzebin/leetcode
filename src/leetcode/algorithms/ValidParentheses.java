package leetcode.algorithms;

import java.util.Stack;

/**
 * @author dzeb
 * @version 1.0
 * @Description 20. Valid Parentheses
 * @createTime 2020/10/26 23:03
 */
public class ValidParentheses {
    public static void main(String[] args) {
        System.out.println(new ValidParentheses().isValid("()"));
        System.out.println(new ValidParentheses().isValid("()[]{}"));
        System.out.println(new ValidParentheses().isValid("(]"));
        System.out.println(new ValidParentheses().isValid("([)]"));
        System.out.println(new ValidParentheses().isValid("{[]}"));
    }
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        char[] chars = s.toCharArray();
        for (char c: chars) {
            if (c == '(' || c == '{' || c == '[') {
                stack.push(c);
                continue;
            }
            if (stack.isEmpty()) {
                return false;
            }
            if (c == ')' && stack.pop() != '(') return false;
            if (c == '}' && stack.pop() != '{') return false;
            if (c == ']' && stack.pop() != '[') return false;
        }
        return stack.isEmpty();
    }
}
