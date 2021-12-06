package leetcode.algorithms;

/**
 * We have two special characters. The first character can be represented by one bit 0. The second character can be represented by two bits (10 or 11).
 *
 * Now given a string represented by several bits. Return whether the last character must be a one-bit character or not. The given string will always end with a zero.
 */
public class IsOneBitCharacter {
    public static void main(String[] args) {
        int[] bits = {1, 1, 1, 0};
        System.out.println(isOneBitCharacter(bits));
    }
    public static boolean isOneBitCharacter(int[] bits) {
        int temp = -1;
        for (int i = 0; i < bits.length - 1; i++) {
            if (temp == -1) {
                if (bits[i] == 0) {
                    continue;
                }
                temp = bits[i];
            } else {
                temp = -1;
            }
        }
        if (temp == -1) {
            return true;
        }
        return false;
    }
}
