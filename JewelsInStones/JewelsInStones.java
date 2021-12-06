package leetcode.algorithms;

/**
 *  You're given strings J representing the types of stones that are jewels, and S representing the stones you have.
 *  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.
 */
public class JewelsInStones {
    public static void main(String[] args) {

    }
    public int numJewelsInStones(String J, String S) {
        int count = 0;
        int arr[] = new int[256];
        for (int i = 0; i < J.length(); i++) {
            arr[J.charAt(i)]++;
        }
        for (int j = 0; j < S.length(); j++) {
            if (arr[S.charAt(j)] != 0) {
                count++;
            }
        }
        return count;
    }
}
