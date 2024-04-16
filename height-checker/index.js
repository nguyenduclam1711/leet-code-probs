/**
 * @param {number[]} heights
 * @return {number}
 */
var heightChecker = function (heights) {
  const sortedHeights = [...heights].sort((a, b) => a - b);

  let res = 0;
  for (let i = 0; i < heights.length; i++) {
    if (sortedHeights[i] !== heights[i]) {
      res++;
    }
  }
  return res;
};
