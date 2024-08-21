/**
 * @param {number[][]} grid
 * @return {number}
 */
var minPathSum = function (grid) {
  const m = grid.length;
  const n = grid[0].length;
  const dp = Array(m);
  for (let i = 0; i < m; i++) {
    dp[i] = Array(n);
  }
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (i === 0 && j === 0) {
        dp[i][j] = grid[i][j];
      } else if (i === 0) {
        dp[i][j] = grid[i][j] + dp[i][j - 1];
      } else if (j === 0) {
        dp[i][j] = grid[i][j] + dp[i - 1][j];
      } else {
        const firstVal = grid[i][j] + dp[i][j - 1];
        const secondVal = grid[i][j] + dp[i - 1][j];
        if (firstVal < secondVal) {
          dp[i][j] = firstVal;
        } else {
          dp[i][j] = secondVal;
        }
      }
    }
  }
  return dp[m - 1][n - 1];
};
