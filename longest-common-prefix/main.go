package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	i := 0
	for {
		if i >= len(strs[0]) {
			return strs[0][:i]
		}
		currStr := strs[0][:i+1]
		allPass := true
		for j := 1; j < len(strs); j++ {
			str := strs[j]
			if i >= len(str) || str[:i+1] != currStr {
				allPass = false
				break
			}
		}
		if allPass {
			i++
		} else {
			break
		}
	}
	return strs[0][:i]
}
