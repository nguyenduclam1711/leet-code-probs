/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  let buyIndex = 0;
  let sellIndex = buyIndex + 1;
  let result = 0;

  while (sellIndex < prices.length) {
    const buyPrice = prices[buyIndex];
    const sellPrice = prices[sellIndex];

    if (sellPrice < buyPrice) {
      buyIndex = sellIndex;
      sellIndex = buyIndex + 1;
      continue;
    }
    const diff = sellPrice - buyPrice;

    if (diff > result) {
      result = diff;
    }
    sellIndex++;
  }
  return result;
};

console.log("Input: prices = [7,1,5,3,6,4]");
console.log(`Output: ${maxProfit([7, 1, 5, 3, 6, 4])}`);
