package leetcode.algorithms.removeDuplicates;

public class RemoveDuplicates {
    public static void main(String[] args) {
        int[] nums = {0,0,1,1,1,2,2,3,3,4};
        int len = removeDuplicates(nums);
        System.out.println("len =" + len);
        for (int i = 0; i < len; i++) {
            System.out.println(nums[i]);
        }
    }

    public static int removeDuplicates(int[] nums) {
        int temp = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] != nums[temp - 1]) {
                nums[temp] = nums[i];
                temp++;
            }
        }
        return temp;
    }
}
