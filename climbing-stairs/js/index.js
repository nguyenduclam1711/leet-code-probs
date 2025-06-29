/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
  // this is simply just calculate fibonacci at n position
  let a = 0;
  let b = 1;

  for (let i = 0; i < n; i++) {
    let tmp = a;
    a = b;
    b += tmp;
  }
  return b;
};

console.log(climbStairs(10));
