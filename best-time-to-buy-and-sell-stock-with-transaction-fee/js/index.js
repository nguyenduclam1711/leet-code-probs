/**
 * @param {number[]} prices
 * @param {number} fee
 * @return {number}
 */
var maxProfit = function(prices, fee) {
  const dp = new Array(prices.length).fill(new Array(2));
  dp[0][0] = 0; // maximum profit when not holding the stock
  dp[0][1] = -prices[0]; // maximum profit when holding the stock

  for (let i = 1; i < prices.length; i++) {
    dp[i][0] = Math.max(prices[i] + dp[i - 1][1] - fee, dp[i - 1][0]);
    dp[i][1] = Math.max(dp[i - 1][0] - prices[i], dp[i - 1][1]);
  }
  return dp[prices.length - 1][0];
};
