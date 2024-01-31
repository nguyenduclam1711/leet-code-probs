/**
 * @param {number[]} arr
 * @return {boolean}
 */
var uniqueOccurrences = function (arr) {
  const m = arr.reduce((acc, n) => {
    acc[n] = (acc[n] || 0) + 1;
    return acc;
  }, {});

  const occur = {};
  for (const n of Object.values(m)) {
    if (!occur[n]) {
      occur[n] = true;
    } else {
      return false;
    }
  }
  return true;
};
