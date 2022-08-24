package exercise_26

import "fmt"

func RemoveDuplicates(nums []int) int {
	for i := 0; i < (len(nums) - 1); {
		if nums[i] == nums[i+1] {
			if i != (len(nums) - 1) {
				nums = append(nums[:i+1], nums[i+2:]...)
			} else {
				nums = nums[:i+1]
			}
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return len(nums)
}
