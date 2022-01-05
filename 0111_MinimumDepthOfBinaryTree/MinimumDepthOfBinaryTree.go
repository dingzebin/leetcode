package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Left)
	if rightDepth == 0 || (leftDepth != 0 && leftDepth < rightDepth) {
		return 1 + leftDepth
	} else {
		return 1 + rightDepth
	}
}
