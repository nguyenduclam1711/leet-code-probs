package designaddandsearchwordsdatastructure

type TrieNode struct {
	data   [26]*TrieNode
	isWord bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	dict := WordDictionary{}
	dict.root = &TrieNode{}

	return dict
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root

	for _, r := range word {
		if r == '.' {
			continue
		}

		idx := r - 'a'
		if curr.data[idx] == nil {
			curr.data[idx] = &TrieNode{}
		}
		curr = curr.data[idx]
	}
	curr.isWord = true
}

func dfs(node *TrieNode, word string) bool {
	if node == nil {
		return false
	}
	if node.isWord && len(word) == 0 {
		return true
	}
	if len(word) == 0 {
		return false
	}

	firstR := word[0]
	if firstR == '.' {
		for _, n := range node.data {
			if n == nil {
				continue
			}
			if dfs(n, word[1:]) {
				return true
			}
		}
		return false
	}

	nodeIdx := firstR - 'a'
	if node.data[nodeIdx] != nil {
		return dfs(node.data[nodeIdx], word[1:])
	}

	return false
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.root, word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
