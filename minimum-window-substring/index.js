/**
 * @param {string} s
 * @param {string} t
 * @return {string}
 */
var minWindow = function(s, t) {
  if (s.length < t.length) {
    return "";
  }
  const mapCountT = {};
  const leftChars = {};
  for (let i = 0; i < t.length; i++) {
    const char = t[i];
    if (!leftChars[char]) {
      leftChars[char] = true;
    }
    mapCountT[char] = (mapCountT[char] ?? 0) + 1;
  }
  let l = 0;
  let r = 0;
  const currCount = {};
  let res = s + " ";

  while (l < s.length && r <= s.length) {
    const isValid = Object.keys(leftChars).length === 0;

    if (!isValid) {
      const char = s[r];
      if (r < s.length) {
        currCount[char] = (currCount[char] ?? 0) + 1;
        if (currCount[char] >= mapCountT[char]) {
          delete leftChars[char];
        }
      }
      r++;
    } else {
      const newRes = s.slice(l, r);
      if (newRes.length < res.length) {
        res = newRes;
      }
      const char = s[l];
      currCount[char]--;
      if (mapCountT[char] > 0 && currCount[char] < mapCountT[char]) {
        leftChars[char] = true;
      }
      l++;
    }
  }
  if (res.length > s.length) {
    return "";
  }
  return res;
};
