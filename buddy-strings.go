package main

func diff(a, b string) []int {
	// diff has the indexes of the different characters
	diff := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff = append(diff, i)
		}
	}

	return diff
}

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		if len(s) == 2 {
			if s[0] == s[1] {
				return true
			}

			return false
		}

		// abab
		if len(s)%2 == 0 && s[:len(s)/2] == s[len(s)/2:] {
			return true
		}

		// aaab
		c := map[rune]struct{}{}
		for _, r := range s {
			c[r] = struct{}{}
		}

		return len(c) < len(s)
	}

	// s can be very long, so we only are going to check the diff
	diff := diff(s, goal)
	if len(diff) == 0 {
		return true
	}

	for _, i := range diff {
		for _, j := range diff {
			if i == j {
				continue
			}

			r := []rune(s)

			r[i], r[j] = r[j], r[i]

			if string(r) == goal {
				return true
			}
		}
	}

	return false
}
