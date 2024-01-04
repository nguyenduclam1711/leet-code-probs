package groupanagrams

type GroupKey [26]int

func GroupAnagrams(strs []string) [][]string {
	groups := make(map[GroupKey][]string)

	for _, str := range strs {
		key := [26]int{}
		for _, r := range str {
			key[r-'a']++
		}
		groups[key] = append(groups[key], str)
	}
	res := make([][]string, len(groups))
	i := 0
	for _, gr := range groups {
		res[i] = gr
		i++
	}
	return res
}
