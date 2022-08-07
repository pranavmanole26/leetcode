package exercise_12

// Convert int to roman number.
// 12. Integer to Roman

func IntToRoman(num int) string {
	roman := ""
	counter := 1
	for num != 0 {
		if counter == 1 {
			roman = getRomanForRange("I", "V", "X", num) + roman
		}
		if counter == 2 {
			roman = getRomanForRange("X", "L", "C", num) + roman
		}
		if counter == 3 {
			roman = getRomanForRange("C", "D", "M", num) + roman
		}
		if counter == 4 {
			n := num % 10
			if n < 4 && n != 0 {
				for n > 0 {
					roman = "M" + roman
					n--
				}
			}
		}
		num = num / 10
		counter++
	}
	return roman
}

func getRomanForRange(fL, mL, lL string, num int) string {
	roman := ""
	n := num % 10
	if n < 4 && n != 0 {
		for n > 0 {
			roman += fL
			n--
		}
	} else if n == 4 {
		roman = fL + mL
	} else if n >= 5 && n < 9 {
		for n > 5 {
			roman += fL
			n--
		}
		roman = mL + roman
	} else if n == 9 {
		roman = fL + lL
	}
	return roman
}
