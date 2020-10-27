package leetcode.algorithms;

import java.util.Arrays;

/**
 * @author dzeb
 * @version 1.0
 * @Description 27. Remove Element
 * @createTime 2020/10/27 22:31
 */
public class RemoveElement {

    public static void main(String[] args) {
        int[] arr = new int[]{1, 1, 3, 4, 4, 5};
        System.out.println(new RemoveElement().removeElement(arr, 4));
        System.out.println(Arrays.toString(arr));
    }

    public int removeElement(int[] nums, int val) {
        int idx = 0;
        int count = 0;
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] != val) {
                nums[idx++] = nums[i];
                count++;
            }
        }
        return count;
    }
}
