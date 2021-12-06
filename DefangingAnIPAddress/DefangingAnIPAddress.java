package leetcode.algorithms;

public class DefangingAnIPAddress {

    public static void main(String[] args) {

    }

    public String defangIPaddr(String address) {
        char[] chars = address.toCharArray();
        StringBuilder sb = new StringBuilder();
        sb.append(chars[0]);
        for(int i = 1; i < chars.length; i++) {
            if (chars[i] == '.') {
                sb.append("[.]");
                sb.append(chars[++i]);
            } else {
                sb.append(chars[i]);
            }
        }
        return sb.toString();
    }
}
