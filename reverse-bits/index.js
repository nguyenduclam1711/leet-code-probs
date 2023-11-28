/**
 * @param {number} n - a positive integer
 * @return {number} - a positive integer
 */
var reverseBits = function(n) {
  let result = 0;
  let pos = 31;

  while (n > 0) {
    result += (n & 1) * Math.pow(2, pos);
    n >>>= 1;
    pos--;
  }
  return result;
};

console.log(reverseBits(43261596));
