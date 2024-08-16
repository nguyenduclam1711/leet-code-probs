/**
 * @param {number[][]} arrays
 * @return {number}
 */
var maxDistance = function(arrays) {
  let firstMax = -Infinity;
  let secondMax = -Infinity;
  let firstMaxIndex = -1;
  let res = 0;

  for (let i = 0; i < arrays.length; i++) {
    const arr = arrays[i];
    const currMax = arr[arr.length - 1];

    if (currMax > firstMax) {
      secondMax = firstMax;
      firstMax = currMax;
      firstMaxIndex = i;
    } else if (currMax > secondMax) {
      secondMax = currMax;
    }
  }
  for (let i = 0; i < arrays.length; i++) {
    const arr = arrays[i];
    let currMax = firstMax;
    if (i === firstMaxIndex) {
      currMax = secondMax;
    }
    const newRes = currMax - arr[0];
    if (newRes > res) {
      res = newRes;
    }
  }
  return res;
};
