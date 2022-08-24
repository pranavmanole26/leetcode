package exercise_15

import (
	"sort"
)

func exist(triplets [][]int, compS []int) bool {
	sort.Ints(compS)
	for _, triplet := range triplets {
		sort.Ints(triplet)
		if compS[0] == triplet[0] && compS[1] == triplet[1] && compS[2] == triplet[2] {
			return true
		}
	}
	return false
}

func ThreeSum(nums []int) [][]int {
	var triplets [][]int
	for i := 0; i < (len(nums) - 2); i++ {
		for j := i + 1; j < (len(nums) - 1); j++ {
			for k := j + 1; k < len(nums); k++ {
				if (nums[i] + nums[j] + nums[k]) == 0 {
					if exist(triplets, []int{nums[i], nums[j], nums[k]}) {
						continue
					}
					var triplet [3]int
					triplet[0] = nums[i]
					triplet[1] = nums[j]
					triplet[2] = nums[k]
					triplets = append(triplets, triplet[:])
				}
			}
		}
	}
	return triplets
}
