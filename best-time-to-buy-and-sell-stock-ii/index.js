/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  prices.push(0);
  let res = 0;
  let i = 0;
  let j = 1;
  let currMax = prices[j];

  while (i < prices.length && j < prices.length) {
    if (prices[i] >= prices[j]) {
      if (currMax > prices[i]) {
        res += currMax - prices[i];
      }
      i = j;
      j = i + 1;
      if (j < prices.length) {
        currMax = prices[j];
      }
    } else {
      if (prices[j] > currMax) {
        currMax = prices[j];
        j++;
      } else {
        res += currMax - prices[i];
        i = j;
        j = i + 1;
        if (j < prices.length) {
          currMax = prices[j];
        }
      }
    }
  }
  prices.pop();
  return res;
};
