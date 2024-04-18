package integertoroman

func convertToRoman(num int, m map[int]string) string {
	if num <= 0 {
		return ""
	}
	if roman, e := m[num]; e {
		return roman
	}
	var res string
	var addon string
	var loop int
	if num < 4 {
		addon = "I"
		loop = num
	} else if num > 5 && num < 9 {
		res = "V"
		addon = "I"
		loop = num - 5
	} else if num > 10 && num < 40 {
		addon = "X"
		loop = num / 10
	} else if num > 50 && num < 90 {
		res = "L"
		addon = "X"
		loop = (num - 50) / 10
	} else if num > 100 && num < 400 {
		addon = "C"
		loop = num / 100
	} else if num > 500 && num < 900 {
		res = "D"
		addon = "C"
		loop = (num - 500) / 100
	} else if num > 1000 && num < 4000 {
		addon = "M"
		loop = num / 1000
	}
	for i := 1; i <= loop; i++ {
		res += addon
	}
	return res
}

func intToRoman(num int) string {
	res := ""
	// s[0] will store thoundsands
	// s[1] will store hundreds
	// s[2] will store dozens
	// s[3] will store units
	s := [4]int{}
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	s[0] = num / 1000 * 1000
	s[1] = (num - s[0]) / 100 * 100
	s[2] = (num - s[0] - s[1]) / 10 * 10
	s[3] = num - s[0] - s[1] - s[2]
	for _, n := range s {
		res += convertToRoman(n, m)
	}
	return res
}
