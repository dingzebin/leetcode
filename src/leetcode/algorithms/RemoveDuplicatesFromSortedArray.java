package leetcode.algorithms;

import java.util.Arrays;

/**
 * @author dzeb
 * @version 1.0
 * @Description 26. Remove Duplicates from Sorted Array
 * @createTime 2020/10/27 22:10
 */
public class RemoveDuplicatesFromSortedArray {

    public static void main(String[] args) {
        int[] arr = new int[]{1, 1, 3, 4, 4, 5};
        System.out.println(new RemoveDuplicatesFromSortedArray().removeDuplicates(arr));
        System.out.println(Arrays.toString(arr));
    }

    public int removeDuplicates(int[] nums) {
        if (nums.length == 0) return 0;
        int count = 1;
        // record the index of last not repeat number
        int idx = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[idx] != nums[i]) {
                nums[++idx] = nums[i];
                count++;
            }
        }
        return count;
    }
}
