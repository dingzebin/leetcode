package leetcode.algorithms;

/**
 * Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.
 *
 * You may return any answer array that satisfies this condition.
 */
public class SortArrayByParity {
    public static void main(String[] args) {
        int[] A = {0, 2};
        new SortArrayByParity().sortArrayByParity(A);
    }
    public int[] sortArrayByParity(int[] A) {
        int first = 0;
        int last = A.length - 1;
        while (first < last) {
            while (A[first] % 2 == 0 && first < last) {
                first++;
            }
            while (A[last] % 2 != 0 && first < last) {
                last--;
            }
            if (first >= last) {
                break;
            }
            A[first] = A[first] ^ A[last];
            A[last] = A[last] ^ A[first];
            A[first] =  A[first] ^ A[last];
            first++;
            last--;
        }
        return A;
    }
}
