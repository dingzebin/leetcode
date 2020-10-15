package leetcode.algorithms;

import java.util.HashMap;
import java.util.Map;

/**
 * 1. Two Sum
 */
public class TwoSum {
    public static void main(String[] args) {

    }
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> temp = new HashMap<>();
        Integer index = null;
        for (int i = 0,len = nums.length; i < len; i++) {
            if ((index = temp.get(target - nums[i])) != null) {
                return new int[]{index, i};
            }
            temp.put(nums[i], i);
        }
        return new int[]{};
    }
}
