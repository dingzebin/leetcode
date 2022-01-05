package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isEqual(root.Left, root.Right)
}

func isEqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil {
		return left.Val == right.Val && isEqual(left.Left, right.Right) && isEqual(left.Right, right.Left)
	}
	return false
}
