package main

func main() {

}

func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[high] {
			if nums[mid] > target && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			return search(nums[low:mid], target) || search(nums[mid+1:high+1], target)
		}
	}
	return false
}
