package dota2senate

func predictPartyVictory(senate string) string {
	stack := []byte(senate)
	i := 0
	allSameParty := false

	for !allSameParty {
		allSameParty = true
		if i >= len(stack) {
			i = 0
		}
		j := i + 1
		for loop := 0; loop < len(stack); loop++ {
			if j >= len(stack) {
				j = 0
			}
			if stack[i] != stack[j] {
				allSameParty = false
				stack = append(stack[:j], stack[j+1:]...)
				break
			}
			j++
		}
		i++
	}
	if stack[0] == 'D' {
		return "Dire"
	}
	return "Radiant"
}
