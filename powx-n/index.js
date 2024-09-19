/**
 * @param {number} x
 * @param {number} n
 * @return {number}
 */
var myPow = function(x, n) {
  if (n === 0) {
    return 1;
  }
  const steps = Math.abs(n);
  const stack = [];
  let i = steps;
  while (i > 1) {
    stack.push(i % 2);
    i = Math.floor(i / 2);
  }

  let res = x;
  while (stack.length > 0) {
    const last = stack[stack.length - 1];
    if (last === 0) {
      res *= res;
    } else {
      res *= res * x;
    }
    stack.pop();
  }
  if (n < 0) {
    return 1 / res;
  }
  return res;
};
