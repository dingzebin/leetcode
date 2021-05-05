package leetcode.algorithms;

import java.util.*;
import java.util.stream.Collector;
import java.util.stream.Collectors;

/**
 * @author dzeb
 * @version 1.0
 * @Description Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
 * Notice that the solution set must not contain duplicate triplets.
 * @createTime 2021/5/5 11:02
 */
public class ThreeSum15 {
    public static void main(String[] args) {
        System.out.println(new ThreeSum15().threeSum(new int[]{0, 0, 0, 0}));
    }
    public List<List<Integer>> threeSum(int[] nums) {
        if(nums.length < 3) return new ArrayList<>();
        Arrays.sort(nums);
        Set<List<Integer>> result = new HashSet<>();
        Set<Integer> set = new HashSet<>();
        for (int i = 0; i < nums.length - 2; i++) {
            int sum = 0 - nums[i];
            set.clear();
            for (int j = i + 1; j < nums.length; j++) {
                if (set.contains(sum - nums[j])) {
                    ArrayList<Integer> list = new ArrayList<>(3);
                    list.add(nums[i]);
                    list.add(sum - nums[j]);
                    list.add(nums[j]);
                    result.add(list);
                } else {
                    set.add(nums[j]);
                }
            }
        }
        return new ArrayList<>(result);
    }
}
