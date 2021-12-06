package leetcode.algorithms;

/**
 * Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).
 *
 * The binary search tree is guaranteed to have unique values.
 */
public class RangeSumBST {
    public static void main(String[] args) {

    }
    public int rangeSumBST(TreeNode root, int L, int R) {
        if (root == null) return 0;
        if (root.val < L) return rangeSumBST(root.right, L, R);
        if (root.val > R) return rangeSumBST(root.left, L, R);
        return root.val + rangeSumBST(root.right, L, R) + rangeSumBST(root.left, L, R);
    }
    private class TreeNode {
        int val;
        TreeNode left;
        TreeNode right;
        TreeNode(int x) { val = x; }
    }
}
