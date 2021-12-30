package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
}

func sortedArrayToBST(nums []int) *TreeNode {
	return covertToTree(nums, 0, len(nums)-1)
}

func covertToTree(nums []int, low int, high int) *TreeNode {
	if low > high {
		return nil
	}
	mid := low + (high-low)/2
	var node = TreeNode{
		Val:   nums[mid],
		Left:  covertToTree(nums, low, mid-1),
		Right: covertToTree(nums, mid+1, high),
	}
	return &node
}
