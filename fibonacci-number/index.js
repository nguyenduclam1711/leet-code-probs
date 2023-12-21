/**
 * @param {number} n
 * @return {number}
 */
var fib = function (n) {
  if (n < 2) {
    return n;
  }
  const arr = [0, 1];

  for (let i = 2; i <= n; i++) {
    arr.push(arr[i - 2] + arr[i - 1]);
  }
  return arr[arr.length - 1];
};
