package leetcode.algorithms;

public class MedianofTwoSortedArrays {
    public static void main(String[] args) {
        int[] nums1 = {1, 2};
        int[] nums2 = {3, 4};
        System.out.println(new MedianofTwoSortedArrays().findMedianSortedArrays(nums1, nums2));
    }

    public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        if (nums1.length <= 0 && nums2.length <= 0) {
            return 0;
        }
        int min = nums1.length > 0 ? nums1[0] : nums2[0];
        int max = nums1.length > 0 ? nums1[0] : nums2[0];
        for(int i = 0; i < nums1.length; i++) {
            min = Math.min(min, nums1[i]);
            max = Math.max(max, nums1[i]);
        }
        for(int i = 0; i < nums2.length; i++) {
            min = Math.min(min, nums2[i]);
            max = Math.max(max, nums2[i]);
        }
        return (min + max) / 2.0;
    }
}
