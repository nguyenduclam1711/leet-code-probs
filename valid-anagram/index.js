/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
  if (s.length !== t.length) {
    return false;
  }
  const dict1 = {};
  const dict2 = {};

  for (let i = 0; i < s.length; i++) {
    dict1[s[i]] = (dict1[s[i]] || 0) + 1;
    dict2[t[i]] = (dict2[t[i]] || 0) + 1;
  }
  if (Object.keys(dict1).length !== Object.keys(dict2).length) {
    return false;
  }
  for (const key of Object.keys(dict1)) {
    if (dict1[key] !== dict2[key]) {
      return false;
    }
  }
  return true;
};
