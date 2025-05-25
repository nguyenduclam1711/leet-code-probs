/**
 * @param {string[]} words
 * @return {number}
 */
var longestPalindrome = function (words) {
  let res = 0
  const wordCount = {};

  for (const w of words) {
    wordCount[w] = (wordCount[w] || 0) + 1;
  }

  const leftPalindromeWords = [];

  for (const [w] of Object.entries(wordCount)) {
    if (wordCount[w] === 0) {
      continue;
    }
    if (w[0] === w[1]) { // this word is already palindrome
      if (wordCount[w] % 2 === 0) {
        res += wordCount[w] * 2;
        wordCount[w] = 0;
      } else if (wordCount[w] > 2) {
        res += (Math.floor(wordCount[w] / 2)) * 4;
        wordCount[w] = 1;
        leftPalindromeWords.push(w);
      } else {
        leftPalindromeWords.push(w);
      }
      continue;
    }
    const otherPalindrome = `${w[1]}${w[0]}`;
    if (wordCount[otherPalindrome] > 0) {
      const added = Math.min(wordCount[w], wordCount[otherPalindrome]);
      res += added * 4;
      wordCount[w] -= added;
      wordCount[otherPalindrome] -= added;
    }
  }
  if (leftPalindromeWords.length > 0) {
    res += 2;
  }
  return res;
};
