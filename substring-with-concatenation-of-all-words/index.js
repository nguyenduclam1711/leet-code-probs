/**
 * @param {string} s
 * @param {string[]} words
 * @return {number[]}
 */
var findSubstring = function (s, words) {
  const hashMap = {};
  const wordLen = words[0].length;
  const windowLen = wordLen * words.length;
  const res = [];

  for (let i = 0; i < words.length; i++) {
    hashMap[words[i]] = (hashMap[words[i]] || 0) + 1;
  }
  for (let i = 0; i <= s.length - windowLen; i++) {
    const newHashMap = { ...hashMap };
    let j = i;
    while (j < i + windowLen) {
      const currWord = s.slice(j, j + wordLen);
      if (!newHashMap[currWord]) {
        break;
      }
      newHashMap[currWord]--;
      j += wordLen;
    }
    if (j === i + windowLen) {
      res.push(i);
    }
  }
  return res;
};
