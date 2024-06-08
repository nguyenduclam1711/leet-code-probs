/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isIsomorphic = function (s, t) {
  const mapping = {};
  const alrMapped = {};

  for (let i = 0; i < s.length; i++) {
    if (!mapping[s[i]]) {
      if (!!alrMapped[t[i]]) {
        return false;
      }
      alrMapped[t[i]] = true;
      mapping[s[i]] = t[i];
    } else if (mapping[s[i]] !== t[i]) {
      return false;
    }
  }
  return true;
};
