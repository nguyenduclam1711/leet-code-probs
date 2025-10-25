package calculatemoneyinleetcodebank

func totalMoney(n int) int {
	total, week, pushMoney := 0, 0, 0
	i := 1

	for i <= n {
		if i%7 == 1 { // monday
			week++
			pushMoney = week
			total += week
		} else {
			pushMoney++
			total += pushMoney
		}
		i++
	}
	return total
}
