/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var findThePrefixCommonArray = function (A, B) {
  const n = A.length;
  const m = Array(n + 1).fill(0);
  const res = [];

  for (let i = 0; i < n; i++) {
    const [a, b] = [A[i], B[i]];
    m[a]++;
    m[b]++;
    let curr = res.length > 0 ? res[i - 1] : 0;

    if (a === b) {
      if (m[a] === 2) {
        curr++;
      }
    } else {
      if (m[a] === 2) {
        curr++;
      }
      if (m[b] === 2) {
        curr++;
      }
    }
    res.push(curr);
  }
  return res;
};
