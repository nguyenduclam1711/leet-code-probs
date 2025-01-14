/**
 * @param {string} s
 * @return {number}
 */
var minimumLength = function (s) {
  const numberOfChars = Array(26).fill(0);
  for (let i = 0; i < s.length; i++) {
    const charIndex = s[i].charCodeAt(0) - 'a'.charCodeAt(0);
    numberOfChars[charIndex]++;
  }
  let res = 0;
  for (const num of numberOfChars) {
    if (num === 0) {
      continue;
    }
    if (num % 2 === 0) {
      res += 2
    } else {
      res++;
    }
  }
  return res;
};
