package leetcode.algorithms;

public class MaxArea {
    public static void main(String[] args) {
        System.out.println(maxArea(new int[]{1,8,6,2,5,4,8,3,7}));
    }
    public static int maxArea(int[] height) {
        int result = 0;
        int temp = 0;
        for(int i = 0; i < height.length; i++) {
            for (int j = i + 1; j < height.length; j++) {
                if (height[j] > height[i]) {
                    temp = height[i] * (j - i);
                } else {
                    temp = height[j] * (j - i);
                }
                if (temp > result) {
                    result = temp;
                }
            }
        }
        return result;
    }

    public int maxArea2(int[] height) {
        int maxarea = 0, l = 0, r = height.length - 1;
        while (l < r) {
            maxarea = Math.max(maxarea, Math.min(height[l], height[r]) * (r - l));
            if (height[l] < height[r])
                l++;
            else
                r--;
        }
        return maxarea;
    }
}
