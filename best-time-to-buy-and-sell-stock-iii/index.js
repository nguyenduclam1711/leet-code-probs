/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
  const firstDp = new Array(prices.length).fill(0);
  const secondDp = Array.from(firstDp);

  let low = prices[0];
  for (let i = 1; i < firstDp.length; i++) {
    const lastHighProfit = firstDp[i - 1];
    const currProfit = prices[i] - low;
    if (currProfit > lastHighProfit) {
      firstDp[i] = currProfit;
    } else {
      firstDp[i] = lastHighProfit;
    }
    if (prices[i] < low) {
      low = prices[i];
    }
  }

  let high = prices[prices.length - 1];
  for (let i = secondDp.length - 2; i >= 0; i--) {
    const lastHighProfit = secondDp[i + 1];
    const currProfit = high - prices[i];
    if (currProfit > lastHighProfit) {
      secondDp[i] = currProfit;
    } else {
      secondDp[i] = lastHighProfit;
    }
    if (prices[i] > high) {
      high = prices[i];
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
