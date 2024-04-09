package optimalpartitionofstring

func partitionString(s string) int {
	exist := [26]bool{}
	res := 1

	for _, c := range s {
		i := c - 'a'
		if exist[i] {
			exist = [26]bool{}
			res++
		}
		exist[i] = true
	}
	return res
}
