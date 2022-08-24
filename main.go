package main

import (
	"fmt"
	"time"

	// exercise_11 "leetcode/exercise_11"
	// exercise_12 "leetcode/exercise_12"
	exercise_15 "leetcode/exercise_15"
	// exercise_20 "leetcode/exercise_20"
	// exercise_26 "leetcode/exercise_26"
)

func main() {
	cur := time.Now()
	// ans := exercise_11.MaxArea(exercise_11.Sl)
	// fmt.Printf("Ans: %t", ans == 49)
	// fmt.Println(exercise_12.IntToRoman(17))
	// fmt.Println(exercise_12.IntToRoman(20))
	// fmt.Println(exercise_12.IntToRoman(33))
	// fmt.Println(exercise_12.IntToRoman(45))
	// fmt.Println(exercise_12.IntToRoman(58))
	// fmt.Println(exercise_12.IntToRoman(69))
	// fmt.Println(exercise_12.IntToRoman(99))
	// fmt.Println(exercise_12.IntToRoman(77))
	// fmt.Println(exercise_12.IntToRoman(1994))
	// fmt.Println(exercise_12.IntToRoman(10))
	// fmt.Println(exercise_12.IntToRoman(20))
	// fmt.Println(exercise_12.IntToRoman(30))
	// fmt.Println(exercise_14.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(exercise_14.LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	// fmt.Println(exercise_14.LongestCommonPrefix([]string{"a", "aca", "accb", "b"}))
	fmt.Println(exercise_15.ThreeSum2([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(exercise_15.ThreeSum2([]int{0, 0, 0}))
	fmt.Println(exercise_15.ThreeSum2([]int{-2, 0, 0, 2, 2}))
	fmt.Println(exercise_15.ThreeSum2([]int{3, 0, -2, -1, 1, 2}))

	// fmt.Println(exercise_5.LongestPalindrome("abb"))
	// fmt.Println(exercise_5.LongestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
	// fmt.Println(exercise_20.IsValid("()"))
	// fmt.Println(exercise_20.IsValid("()[]{}"))
	// fmt.Println(exercise_20.IsValid("(]"))

	// fmt.Println(exercise_26.RemoveDuplicates([]int{1, 1, 2}))
	// fmt.Println(exercise_26.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}))
	// fmt.Println(exercise_27.RemoveElement([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4}, 1))

	fmt.Printf("Time taken: %d", time.Since(cur))
}
