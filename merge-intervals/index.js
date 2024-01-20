/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function(intervals) {
  if (intervals.length <= 1) {
    return intervals;
  }
  const sortedIntervals = intervals.sort((a, b) => a[0] - b[0]);
  const result = [sortedIntervals[0]];
  let i = 0;
  let j = 1;

  while (j < sortedIntervals.length) {
    const resInterval = result[i];
    const currInterval = sortedIntervals[j];

    if (currInterval[0] > resInterval[1]) {
      result.push(currInterval);
      i++;
    } else if (currInterval[1] > resInterval[1]) {
      resInterval[1] = currInterval[1];
    }
    j++;
  }
  return result;
};
