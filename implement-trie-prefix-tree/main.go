package implementtrieprefixtree

type TrieNode struct {
	data   [26]*TrieNode
	isWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	t := Trie{}
	t.root = &TrieNode{}
	return t
}

func (this *Trie) Insert(word string) {
	curr := this.root

	for _, r := range word {
		idx := r - 'a'
		if curr.data[idx] == nil {
			curr.data[idx] = &TrieNode{}
		}
		curr = curr.data[idx]
	}
	curr.isWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for _, r := range word {
		idx := r - 'a'

		if curr.data[idx] == nil {
			return false
		}
		curr = curr.data[idx]
	}
	return curr.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root

	for _, r := range prefix {
		idx := r - 'a'
		if curr.data[idx] == nil {
			return false
		}
		curr = curr.data[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
