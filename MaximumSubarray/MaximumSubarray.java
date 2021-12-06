package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 53. Maximum Subarray
 * @createTime 2020/10/16 23:40
 */
public class MaximumSubarray {
    public static void main(String[] args) {
        System.out.println(new MaximumSubarray().maxSubArray(new int[]{-2,1,-3,4,-1,2,1,-5,4}));
    }

    public int maxSubArray(int[] nums) {
        int sum = nums[0];
        int res = sum;

        for (int i = 1; i < nums.length; i++) {
            sum = Math.max(sum + nums[i], nums[i]);
            res = Math.max(sum, res);
        }
        return res;
    }
}
