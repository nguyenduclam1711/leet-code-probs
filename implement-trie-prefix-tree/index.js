class Trie {
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
  insert(word) {
    let curr = this.root;
    for (const s of word) {
      let idx = s.charCodeAt(0) - "a".charCodeAt(0);

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
   * @param {string} word
   * @return {boolean}
   */
  search(word) {
    let curr = this.root;

    for (const s of word) {
      let idx = s.charCodeAt(0) - "a".charCodeAt(0);

      if (!curr.data[idx]) {
        return false;
      }
      curr = curr.data[idx];
    }
    return curr.isWord;
  }
  /**
   * @param {string} prefix
   * @return {boolean}
   */
  startsWith(prefix) {
    let curr = this.root;

    for (const s of prefix) {
      let idx = s.charCodeAt(0) - "a".charCodeAt(0);

      if (!curr.data[idx]) {
        return false;
      }
      curr = curr.data[idx];
    }
    return true;
  }
}

/**
 * Your Trie object will be instantiated and called as such:
 * var obj = new Trie()
 * obj.insert(word)
 * var param_2 = obj.search(word)
 * var param_3 = obj.startsWith(prefix)
 */
