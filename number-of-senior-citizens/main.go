package numberofseniorcitizens

import "strconv"

func countSeniors(details []string) int {
	res := 0
	for _, detail := range details {
		age, _ := strconv.Atoi(detail[11:13])
		if age > 60 {
			res++
		}
	}
	return res
}
