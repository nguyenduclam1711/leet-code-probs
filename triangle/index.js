/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function (triangle) {
  const dp = Array(triangle.length);
  dp[0] = triangle[0];

  for (let i = 1; i < dp.length; i++) {
    dp[i] = Array(i + 1);
    for (let j = 0; j < dp[i].length; j++) {
      if (j === 0) {
        dp[i][j] = triangle[i][j] + dp[i - 1][0];
      } else if (j === i) {
        dp[i][j] = triangle[i][j] + dp[i - 1][dp[i - 1].length - 1];
      } else {
        const firstVal = triangle[i][j] + dp[i - 1][j - 1];
        const secondVal = triangle[i][j] + dp[i - 1][j];
        if (firstVal < secondVal) {
          dp[i][j] = firstVal;
        } else {
          dp[i][j] = secondVal;
        }
      }
    }
  }
  return Math.min(...dp[dp.length - 1]);
};
