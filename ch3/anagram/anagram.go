package anagram

func isAnagram(s1, s2 string) bool {
	letters1 := make(map[rune]int)
	letters2 := make(map[rune]int)

	if len(letters1) != len(letters2) {
		return false
	}

	for _, r := range s1 {
		letters1[r]++
	}

	for _, r := range s2 {
		letters2[r]++
	}

	for k, v := range letters1 {
		if letters2[k] != v {
			return false
		}
	}
	return true
}
