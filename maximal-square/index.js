/**
 * @param {character[][]} matrix
 * @return {number}
 */
var maximalSquare = function(matrix) {
  const dp = Array(matrix.length);
  for (let i = 0; i < matrix.length; i++) {
    dp[i] = Array(matrix[0].length);
  }
  let edgeLen = 0;
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      const squareVal = matrix[i][j];

      if (squareVal === "0") {
        dp[i][j] = 0;
      } else if (i === 0 || j === 0) {
        dp[i][j] = 1;
      } else if (squareVal === "1") {
        const leftVal = dp[i][j - 1];
        const upVal = dp[i - 1][j];
        const diagnolVal = dp[i - 1][j - 1];
        const minVal = Math.min(leftVal, upVal, diagnolVal);

        if (minVal === 0) {
          dp[i][j] = 1;
        } else {
          dp[i][j] = minVal + 1;
        }
      }
      if (dp[i][j] > edgeLen) {
        edgeLen = dp[i][j];
      }
    }
  }
  return edgeLen * edgeLen;
};
