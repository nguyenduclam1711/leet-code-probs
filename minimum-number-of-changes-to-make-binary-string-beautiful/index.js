/**
 * @param {string} s
 * @return {number}
 */
var minChanges = function(s) {
  let res = 0;
  for (let i = 0; i < s.length; i += 2) {
    const first = s[i];
    const second = s[i + 1];
    if (first !== second) {
      res++;
    }
  }
  return res;
};
