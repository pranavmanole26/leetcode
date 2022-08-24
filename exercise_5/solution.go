package exercise_5

func LongestPalindrome(s string) string {
	pal := ""

	if len(s) == 1 {
		return s
	}
	for 0 < len(s) {
		count := 1
		for j := 1; count <= len(s); j++ {
			tempPal := ""
			if count == len(s) {
				if s[:j] == rev(s[j-1:]) {
					tempPal = s[:j-1] + s[j-1:]
					if len(pal) < len(tempPal) {
						pal = tempPal
					}
				}

				break
			}
			if s[:j] == rev(s[j-1:count]) {
				tempPal = s[:j-1] + s[j-1:count]
				if len(pal) < len(tempPal) {
					pal = tempPal
				}
			}
			count += 1
			if count == len(s) {
				if s[:j] == rev(s[j:]) {
					tempPal = s[:j] + s[j:]
					if len(pal) < len(tempPal) {
						pal = tempPal
					}
				}
				break
			}
			if s[:j] == rev(s[j:count]) {
				tempPal = s[:j] + s[j:count]
				if len(pal) < len(tempPal) {
					pal = tempPal
				}
			}
			count += 1
		}

		s = s[1:]
		if len(s) < len(pal) {
			return pal
		}
	}
	return pal
}

func rev(str string) string {
	r := ""
	for _, s := range str {
		r = string(s) + r
	}
	return r
}
