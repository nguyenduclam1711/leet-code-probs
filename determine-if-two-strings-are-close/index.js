/**
 * @param {string} word
 * @return {Array<number>}
 */
function getMapChars(word) {
  const arr = Array(26).fill(0);
  for (const c of word) {
    const idx = c.charCodeAt(0) - "a".charCodeAt(0);
    arr[idx]++;
  }
  return arr;
}

/**
 * @param {string} word1
 * @param {string} word2
 * @return {boolean}
 */
var closeStrings = function(word1, word2) {
  if (word1.length !== word2.length) {
    return false;
  }
  const mapChars1 = getMapChars(word1);
  const mapChars2 = getMapChars(word2);

  for (let i = 0; i < mapChars1.length; i++) {
    if (mapChars1[i] > 0 && mapChars2[i] === 0) {
      return false;
    }
  }
  mapChars1.sort((a, b) => a - b);
  mapChars2.sort((a, b) => a - b);
  for (let i = 0; i < mapChars1.length; i++) {
    if (mapChars1[i] !== mapChars2[i]) {
      return false;
    }
  }
  return true;
};
