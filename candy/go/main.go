package candy

// using greedy to solve this
// find relative smallest position and then set the candy to that position to 1 and then from that position, distribute the rest of candies
func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	distributions := make([]int, len(ratings))
	initDistributionPos := []int{}

	for i, r := range ratings {
		if i == 0 {
			if ratings[i+1] >= r {
				distributions[i] = 1
				initDistributionPos = append(initDistributionPos, i)
			}
		} else if i == len(ratings)-1 {
			if ratings[i-1] >= r {
				distributions[i] = 1
				initDistributionPos = append(initDistributionPos, i)
			}
		} else {
			if ratings[i-1] >= r && ratings[i+1] >= r {
				distributions[i] = 1
				initDistributionPos = append(initDistributionPos, i)
			}
		}
	}
	for _, pos := range initDistributionPos {
		// move backward
		for i := pos - 1; i >= 0; i-- {
			if ratings[i] <= ratings[i+1] {
				break
			}
			distributions[i] = max(distributions[i], distributions[i+1]+1)
		}
		// move forward
		for i := pos + 1; i < len(ratings); i++ {
			if ratings[i] <= ratings[i-1] {
				break
			}
			distributions[i] = max(distributions[i], distributions[i-1]+1)
		}
	}
	res := 0
	for _, d := range distributions {
		res += d
	}
	return res
}
