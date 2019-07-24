package leetcode.algorithms.maxSubArray;

/**
 * @description:
 * @author: dzeb
 * @date: 2019/7/24 15:09
 **/
public class MaxSubArray {
    public static void main(String[] args) {
        int[] arr = {-2,1,-3,4,-1,2,1,-5,4};
        System.out.println(maxSubArray(arr));
    }

    public static int maxSubArray(int[] nums) {
        int maxSum = nums[0];
        int temp = nums[0];
        for(int i = 1; i < nums.length; i++) {
            temp = temp > 0 ? temp + nums[i] : nums[i];
            if (temp > maxSum) {
                maxSum = temp;
            }
        }
        return maxSum;
    }
}
