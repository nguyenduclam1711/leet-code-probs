/**
 * @param {number} a
 * @param {number} b
 * @return {number}
 */
var getSum = function(a, b) {
  let res = a ^ b;
  let carry = (a & b) << 1;

  while (carry !== 0) {
    const tmp = res;
    res ^= carry;
    carry = (tmp & carry) << 1;
  }
  return res;
};

console.log(getSum(1, 2));
