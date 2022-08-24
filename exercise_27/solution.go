package exercise_27

import "fmt"

func RemoveElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if val == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	fmt.Println(nums)
	return len(nums)
}
