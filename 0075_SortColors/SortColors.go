package main

func main() {

}

func sortColors(nums []int) {
	red, white, blue := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			red++
		} else if nums[i] == 1 {
			white++
		} else {
			blue++
		}
	}
	i := 0
	for ; red > 0; red-- {
		nums[i] = 0
		i++
	}
	for ; white > 0; white-- {
		nums[i] = 1
		i++
	}
	for ; blue > 0; blue-- {
		nums[i] = 2
		i++
	}
}
