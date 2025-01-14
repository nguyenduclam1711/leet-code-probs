/**
 * @param {string[]} words
 * @return {string[]}
 */
var stringMatching = function (words) {
  const res = [];
  for (const word1 of words) {
    for (const word2 of words) {
      if (word1 === word2) {
        continue;
      }
      if (word2.includes(word1)) {
        res.push(word1);
        break;
      }
    }
  }
  return res;
};
