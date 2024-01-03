/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  if (s.length <= 1) {
    return s.length;
  }
  let l = 0;
  let r = 1;
  let maxLen = 1;
  const hashMap = {
    [s[l]]: true,
  };

  while (l < s.length && r < s.length) {
    if (l === r) {
      r = l + 1;
      hashMap[s[l]] = true;
      continue;
    }
    if (!hashMap[s[r]]) {
      hashMap[s[r]] = true;
      const currLen = r - l + 1;

      if (currLen > maxLen) {
        maxLen = currLen;
      }
      r++;
    } else {
      hashMap[s[l]] = false;
      l++;
    }
  }
  return maxLen;
};
