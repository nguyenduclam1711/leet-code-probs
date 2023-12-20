package numberof1bits

func HammingWeight(num uint32) int {
	w := 0

	for num != 0 {
		if num&1 == 1 {
			w++
		}
		num = num >> 1
	}
	return w
}
