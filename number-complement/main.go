package numbercomplement

func findComplement(num int) int {
	temp := num
	mask := 1

	for temp > 0 {
		mask = (mask << 1) | 1
		temp = temp >> 1
	}
	return num ^ mask
}
