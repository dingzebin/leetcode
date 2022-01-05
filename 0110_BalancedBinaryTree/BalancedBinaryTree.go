package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if isBalancedTree(root) != -1 {
		return true
	}
	return false
}

func isBalancedTree(node *TreeNode) int {
	if node == nil {
		return 0
	}
	max, min := sort(isBalancedTree(node.Left), isBalancedTree(node.Right))
	if max == -1 || min == -1 || max-min > 1 {
		return -1
	}
	return 1 + max
}

func sort(x, y int) (int, int) {
	if x > y {
		return x, y
	}
	return y, x
}
