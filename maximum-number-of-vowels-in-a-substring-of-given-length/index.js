/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var maxVowels = function(s, k) {
  const vowels = {
    a: true,
    e: true,
    i: true,
    o: true,
    u: true,
  };
  let l = 0;
  let r = 0;
  let currNumberOfVowels = 0;
  let res = 0;

  while (l < s.length && r < s.length) {
    const windowLen = r - l + 1;

    if (vowels[s[r]]) {
      currNumberOfVowels++;
    }
    if (windowLen > k) {
      if (vowels[s[l]]) {
        currNumberOfVowels--;
      }
      l++;
    }
    if (currNumberOfVowels > res) {
      res = currNumberOfVowels;
    }
    r++;
  }
  return res;
};
