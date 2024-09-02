package combinations

func combine(n int, k int) [][]int {
	res := [][]int{}
	arr := []int{}
	var recurse func(i int)
	recurse = func(i int) {
		if len(arr) == k {
			newArr := make([]int, len(arr))
			copy(newArr, arr)
			res = append(res, newArr)
			return
		}
		for j := i + 1; j <= n; j++ {
			arr = append(arr, j)
			recurse(j)
			arr = arr[:len(arr)-1]
		}
	}
	recurse(0)
	return res
}
