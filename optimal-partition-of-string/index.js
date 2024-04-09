/**
 * @param {string} s
 * @return {number}
 */
var partitionString = function (s) {
  const initArr = Array(26).fill(false);
  let exist = [...initArr];
  let res = 1;
  for (const c of s) {
    const i = c.charCodeAt(0) - "a".charCodeAt(0);
    if (exist[i]) {
      exist = [...initArr];
      res++;
    }
    exist[i] = true;
  }
  return res;
};
