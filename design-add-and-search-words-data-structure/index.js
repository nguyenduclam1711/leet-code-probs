class WordDictionary {
  constructor() {
    this.root = {
      data: Array(26).fill(null),
      isWord: false,
    };
  }

  /**
   * @param {string} word
   * @return {void}
   */
  addWord(word) {
    let curr = this.root;

    for (const s of word) {
      const idx = s.charCodeAt(0) - "a".charCodeAt(0);

      if (!curr.data[idx]) {
        curr.data[idx] = {
          data: Array(26).fill(null),
          isWord: false,
        };
      }
      curr = curr.data[idx];
    }
    curr.isWord = true;
  }

  /**
   * @param {object | null} node
   * @param {string} word
   * @return {boolean}
   */
  dfs(node, word) {
    if (!node) {
      return false;
    }
    if (node.isWord && word.length === 0) {
      return true;
    }
    if (word.length === 0) {
      return false;
    }

    const firstS = word[0];
    if (firstS === ".") {
      for (const n of node.data) {
        if (!n) {
          continue;
        }
        if (this.dfs(n, word.slice(1))) {
          return true;
        }
      }
    }

    const nodeIdx = firstS.charCodeAt(0) - "a".charCodeAt(0);
    if (node.data[nodeIdx]) {
      return this.dfs(node.data[nodeIdx], word.slice(1));
    }

    return false;
  }

  /**
   * @param {string} word
   * @return {boolean}
   */
  search(word) {
    return this.dfs(this.root, word);
  }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * var obj = new WordDictionary()
 * obj.addWord(word)
 * var param_2 = obj.search(word)
 */
