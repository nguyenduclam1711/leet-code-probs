/**
 * @param {number} n
 * @param {number} k
 * @return {number[][]}
 */
var combine = function(n, k) {
  const res = [];
  const arr = [];

  function recurse(i) {
    if (arr.length === k) {
      res.push([...arr]);
      return;
    }
    for (let j = i + 1; j <= n; j++) {
      arr.push(j);
      recurse(j);
      arr.pop();
    }
  }
  recurse(0);
  return res;
};
