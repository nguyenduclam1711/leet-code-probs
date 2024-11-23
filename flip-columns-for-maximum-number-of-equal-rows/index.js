function count(arr, countPattern) {
  let currN = -1;
  let currNCount = 1;
  let patternKey = "";

  for (const n of arr) {
    if (n !== currN) {
      patternKey += `,${currNCount}`;
      currN = n;
      currNCount = 1;
    } else {
      currNCount++;
    }
  }
  patternKey += `,${currNCount}`;
  countPattern.set(patternKey, (countPattern.get(patternKey) || 0) + 1);
}

/**
 * @param {number[][]} matrix
 * @return {number}
 */
var maxEqualRowsAfterFlips = function (matrix) {
  const countPattern = new Map();

  for (const row of matrix) {
    count(row, countPattern);
  }
  return Math.max(...countPattern.values());
};
