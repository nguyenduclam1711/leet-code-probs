/**
 * @param {number[][]} matrix
 * @param {number} target
 * @return {boolean}
 */
var searchMatrix = function(matrix, target) {
  let res = false;

  for (let i = 0; i < matrix.length; i++) {
    const row = matrix[i];
    const last = row[row.length - 1];

    if (target > last) {
      continue;
    }
    let l = 0;
    let r = row.length;
    while (l < r) {
      const m = Math.floor((l + r) / 2);
      if (target === row[m]) {
        res = true;
        break;
      } else if (target > row[m]) {
        l = m + 1;
      } else {
        r = m;
      }
    }
    break;
  }
  return res;
};
