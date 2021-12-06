package leetcode.algorithms;

import java.util.Arrays;

/**
 * @author dzeb
 * @version 1.0
 * @Description
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 * Write the code that will take a string and make this conversion given a number of rows:
 *
 * string convert(string s, int numRows);
 * @createTime 2021/5/4 11:29
 */
public class ZigZagConversion6 {
    public static void main(String[] args) {
        System.out.println("PAYPALISHIRING");
        System.out.println(new ZigZagConversion6().convert("PAYPALISHIRING", 4));
        System.out.println("PINALSIGYAHRPI");
    }
    public String convert(String s, int numRows) {
        int len = s.length();
        if (numRows == 1 || len <= 2) return s;
        char[][] array = new char[numRows][len / numRows * 2 + 1];
        int row = 0;
        int column = 0;
        boolean down = true;
        for (int i = 0; i < len; i++) {
            array[row][column] = s.charAt(i);
            if (i != 0 && (row == 0 || row == numRows - 1)) {
                down = !down;
                column++;
            }
            row += down ? 1 : -1;
        }
        StringBuilder sb = new StringBuilder();
        for(int i = 0; i < numRows; i++) {
            for (int j = 0; j < array[i].length; j++) {
                if (array[i][j] != '\0') {
                    sb.append(array[i][j]);
                }
            }
        }
        return sb.toString();
    }
}
