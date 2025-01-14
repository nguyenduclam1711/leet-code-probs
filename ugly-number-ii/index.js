/**
 * @param {number} n
 * @return {number}
 */
var nthUglyNumber = function (n) {
  const ugly = Array(n);
  ugly[0] = 1;
  let [index1, index2, index3] = [0, 0, 0];
  let [factor1, factor2, factor3] = [2, 3, 5];

  for (let i = 1; i < n; i++) {
    const nextUgly = Math.min(factor1, factor2, factor3);
    ugly[i] = nextUgly;
    if (nextUgly === factor1) {
      index1++;
      factor1 = 2 * ugly[index1];
    }
    if (nextUgly === factor2) {
      index2++;
      factor2 = 3 * ugly[index2];
    }
    if (nextUgly === factor3) {
      index3++;
      factor3 = 5 * ugly[index3];
    }
  }
  return ugly[n - 1];
};
