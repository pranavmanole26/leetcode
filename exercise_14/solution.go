package exercise_14

import "math"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	longestCPref := strs[0]

	for n := 0; n < (len(strs) - 1); n++ {
		if len(strs[n]) == 0 || len(strs[n+1]) == 0 {
			return ""
		}

		pref := ""
		shortS := int(math.Min(float64(len(strs[n])), float64(len(strs[n+1]))))

		for i := 0; i < shortS; i++ {
			if strs[n][i] == strs[n+1][i] {
				pref += string(strs[n][i])
			} else {
				break
			}
		}
		if len(pref) == 0 {
			return ""
		}
		if len(longestCPref) >= len(pref) {
			longestCPref = pref
		}
	}
	return longestCPref
}
