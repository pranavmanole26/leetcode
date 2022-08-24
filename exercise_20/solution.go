package exercise_20

func IsValid(s string) bool {

	t := make([]rune, 0)
	for _, bracket := range s {
		if len(t) == 0 {
			t = append(t, bracket)
		} else {
			if string(bracket) == ")" && string(t[len(t)-1]) == "(" ||
				string(bracket) == "}" && string(t[len(t)-1]) == "{" ||
				string(bracket) == "]" && string(t[len(t)-1]) == "[" {
				t = t[:len(t)-1]
			} else {
				t = append(t, bracket)
			}
		}
	}
	return len(t) == 0
}
