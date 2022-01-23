package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1
	for start < end {
		tmp := numbers[start] + numbers[end]
		if tmp == target {
			return []int{start + 1, end + 1}
		} else if tmp < target {
			start++
		} else {
			end--
		}
	}
	return []int{}
}
