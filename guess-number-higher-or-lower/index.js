/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * var guess = function(num) {}
 */

/**
 * @param {number} n
 * @return {number}
 */
var guessNumber = function(n) {
  let l = 1;
  let r = n;
  while (l <= r) {
    const m = Math.floor((l + r) / 2);
    const guessRes = guess(m);
    if (guessRes == 0) {
      return m;
    } else if (guessRes > 0) {
      l = m + 1;
    } else {
      r = m;
    }
  }
  return -1;
};
