/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  const firstDp = new Array(prices.length).fill(0);
  const secondDp = Array.from(firstDp);
  let low = prices[0];
  let high = prices[prices.length - 1];

  for (let i = 1; i < prices.length; i++) {
    firstDp[i] = Math.max(prices[i] - low, firstDp[i - 1]);
    if (prices[i] < low) {
      low = prices[i];
    }
    const j = prices.length - i - 1;
    secondDp[j] = Math.max(high - prices[j], secondDp[j + 1]);
    if (prices[j] > high) {
      high = prices[j];
    }
  }

  let res = 0;
  for (let i = 0; i < firstDp.length; i++) {
    let currRes = firstDp[i];
    if (i < firstDp.length - 1) {
      currRes += secondDp[i + 1];
    }
    if (currRes > res) {
      res = currRes;
    }
  }
  return res;
};
