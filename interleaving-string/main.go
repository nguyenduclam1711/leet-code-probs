package interleavingstring

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s3) == 0 {
		return true
	}
	q := [][]string{
		{s1, s2},
	}
	compareStr := ""
	for i := 0; i < len(s3); i++ {
		c := s3[i]
		newQ := [][]string{}
		existedS1 := map[string]bool{}
		existedS2 := map[string]bool{}

		for _, strs := range q {
			currS1, currS2 := strs[0], strs[1]
			if len(currS1) == 0 && len(currS2) == 0 {
				continue
			}
			if len(currS1) > 0 && currS1[0] == c {
				nextS1 := currS1[1:]
				if !existedS2[nextS1] || !existedS1[currS2] {
					newQ = append(newQ, []string{nextS1, currS2})
					existedS2[nextS1] = true
					existedS1[currS2] = true
				}
			}
			if len(currS2) > 0 && currS2[0] == c {
				nextS2 := currS2[1:]
				if !existedS2[nextS2] || !existedS1[currS1] {
					newQ = append(newQ, []string{currS1, nextS2})
					existedS2[nextS2] = true
					existedS1[currS1] = true
				}
			}
		}
		if len(newQ) > 0 {
			compareStr += string(c)
		}
		q = newQ
	}
	return compareStr == s3
}
