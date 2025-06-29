/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
  const max = amount + 1;
  const dp = Array(max);
  const sortedCoins = coins.sort((a, b) => a - b);
  dp[0] = 0;
  for (let i = 1; i < dp.length; i++) {
    dp[i] = Infinity;
    for (const coin of sortedCoins) {
      if (i === coin) {
        dp[i] = 1;
        break;
      }
      if (i > coin) {
        dp[i] = Math.min(dp[i], 1 + dp[i - coin]);
      } else {
        break;
      }
    }
  }
  if (dp[amount] === Infinity) {
    return -1;
  }
  return dp[amount];
};
