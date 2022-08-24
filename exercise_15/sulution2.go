package exercise_15

import (
	"fmt"
	"sort"
)

func ThreeSum2(nums []int) [][]int {
	var triplets [][]int

	sort.Ints(nums)

	left := 0
	right := len(nums) - 1
	sum := 0

	for {
		sum = nums[left] + nums[right]
		if sum < 0 {
			for j := right - 1; nums[j] >= 0; j-- {
				if sum+nums[j] == 0 {
					if len(triplets) > 0 && !(triplets[len(triplets)-1][0] == nums[left] && triplets[len(triplets)-1][1] == nums[right]) {
						triplets = append(triplets, []int{nums[left], nums[right], nums[j]})
					} else if len(triplets) == 0 {
						triplets = append(triplets, []int{nums[left], nums[right], nums[j]})
					}
					break
				}
			}
		} else {
			for j := left + 1; nums[j] <= 0; j++ {
				if sum+nums[j] == 0 {
					if len(triplets) > 0 && !(triplets[len(triplets)-1][0] == nums[left] && triplets[len(triplets)-1][1] == nums[right]) {
						triplets = append(triplets, []int{nums[left], nums[right], nums[j]})
					} else if len(triplets) == 0 {
						triplets = append(triplets, []int{nums[left], nums[right], nums[j]})
					}
					break
				}
			}
		}
		if sum < 0 {
			left++
		} else if sum == 0 {

		} else {
			right--
		}

		if left > right || right-left < 2 {
			break
		}
	}
	fmt.Println(triplets)
	return triplets
}
