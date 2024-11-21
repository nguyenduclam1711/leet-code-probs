/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var takeCharacters = function(s, k) {
  const countChars = [0, 0, 0];
  for (let i = 0; i < s.length; i++) {
    countChars[s[i].charCodeAt(0) - "a".charCodeAt(0)]++;
  }
  if (countChars[0] < k || countChars[1] < k || countChars[2] < k) {
    return -1;
  }
  let res = s.length;
  let l = 0;
  let r = 0;
  let windowSize = 0;
  const currCountChars = [0, 0, 0];
  while (r < s.length) {
    currCountChars[s[r].charCodeAt(0) - "a".charCodeAt(0)]++;
    windowSize = r - l + 1;
    const countLeftA = countChars[0] - currCountChars[0];
    const countLeftB = countChars[1] - currCountChars[1];
    const countLeftC = countChars[2] - currCountChars[2];
    if (countLeftA >= k && countLeftB >= k && countLeftC >= k) {
      const newRes = s.length - windowSize;
      if (newRes < res) {
        res = newRes;
      }
    } else {
      currCountChars[s[l].charCodeAt(0) - "a".charCodeAt(0)]--;
      l++;
    }
    r++;
  }
  return res;
};
