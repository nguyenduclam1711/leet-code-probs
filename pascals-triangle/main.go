package pascalstriangle

func Generate(numRows int) [][]int {
	result := [][]int{}
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			result = append(result, []int{1})
			continue
		}
		if i == 2 {
			result = append(result, []int{1, 1})
			continue
		}
		row := make([]int, i)
		l, r := 0, 1
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				row[j] = 1
			} else {
				row[j] = result[i-2][l] + result[i-2][r]
				l++
				r++
			}
		}
		result = append(result, row)
	}
	return result
}
