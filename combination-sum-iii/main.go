package combinationsumiii

func dfs(num int, k int, n int, total int, nums *[]int, res *[][]int) {
	if total > n {
		return
	}
	if len(*nums) == k {
		if total == n {
			*res = append(*res, append([]int{}, *nums...))
		}
		return
	}
	for i := num + 1; i < 10; i++ {
		*nums = append(*nums, i)
		dfs(i, k, n, total+i, nums, res)
		*nums = (*nums)[:len(*nums)-1]
	}
}

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	nums := []int{}
	dfs(0, k, n, 0, &nums, &res)
	return res
}
