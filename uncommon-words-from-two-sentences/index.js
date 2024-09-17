/**
 * @param {string} s1
 * @param {string} s2
 * @return {string[]}
 */
var uncommonFromSentences = function(s1, s2) {
  const countWords = {};
  function count(s) {
    if (!countWords[s]) {
      countWords[s] = 1;
    } else {
      countWords[s]++;
    }
  }
  s1.split(" ").forEach(count);
  s2.split(" ").forEach(count);

  const res = [];
  Object.keys(countWords).forEach((s) => {
    if (countWords[s] === 1) {
      res.push(s);
    }
  });
  return res;
};
