package leetcode.algorithms.flippingAnImage;

import java.util.Arrays;

public class FlippingAnImage {
    public static void main(String[] args) {
        int[][] a = {{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}};
        flipAndInvertImage(a);
        for(int i = 0; i < a.length; i++) {
            System.out.println(Arrays.toString(a[i]));
        }
    }

    /**
     * [[1,1,0],
     *  [1,0,1],
     *  [0,0,0]]
     *
     *  [[1,0,0],
     *  [0,1,0],
     *  [1,1,1]]
     * @param A
     * @return
     */
    public static int[][] flipAndInvertImage(int[][] A) {
        for (int i = 0; i < A.length; i++) {
            for (int j = 0, len = A[i].length >> 1; j <= len; j++) {
                int k = A[i].length - j - 1;
                if (j <= k) {
                    A[i][j] ^= 1;
                    if (j < k) {
                        A[i][k] ^= 1;
                        A[i][j] ^= A[i][k];
                        A[i][k] ^= A[i][j];
                        A[i][j] ^= A[i][k];
                    }
                }
            }
        }
        return A;
    }
}
