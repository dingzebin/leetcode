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
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 {
		return 1 + rightDepth
	} else if rightDepth == 0 {
		return 1 + leftDepth
	} else {
		return 1 + min(leftDepth, rightDepth)
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
