/**
 * @param {string} pattern
 * @param {string} s
 * @return {boolean}
 */
var wordPattern = function(pattern, s) {
  const splitedS = s.split(" ");
  if (splitedS.length !== pattern.length) {
    return false;
  }
  const m = new Map();
  const alrMapped = new Map();

  for (let i = 0; i < pattern.length; i++) {
    if (!m.has(pattern[i])) {
      if (alrMapped.has(splitedS[i])) {
        return false;
      }
      m.set(pattern[i], splitedS[i]);
      alrMapped.set(splitedS[i], true);
    } else if (m.get(pattern[i]) !== splitedS[i]) {
      return false;
    }
  }
  return true;
};
