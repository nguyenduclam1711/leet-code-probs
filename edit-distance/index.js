/**
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
var minDistance = function (word1, word2) {
  const dp = Array(word1.length + 1);
  for (let i = 0; i < dp.length; i++) {
    dp[i] = Array(word2.length + 1).fill(0);
    if (i === 0) {
      for (let j = 0; j < dp[i].length; j++) {
        dp[i][j] = j;
      }
    }
    dp[i][0] = i;
  }
  for (let i = 1; i < dp.length; i++) {
    for (let j = 1; j < dp[i].length; j++) {
      if (word1[i - 1] === word2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = Math.min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]) + 1;
      }
    }
  }
  return dp[word1.length][word2.length];
};
