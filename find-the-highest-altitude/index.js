/**
 * @param {number[]} gain
 * @return {number}
 */
var largestAltitude = function (gain) {
  let curr = 0;
  let res = 0;

  for (const n of gain) {
    curr += n;
    if (curr > res) {
      res = curr;
    }
  }
  return res;
};
